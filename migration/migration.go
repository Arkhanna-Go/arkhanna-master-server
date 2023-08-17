package migration

// TODO: Create a new file (or package) for handling migration information saved on DB
// CREATE TABLE migration (
// 	index 		INT,
// 	file  		VARCHAR(50),
// 	filesha256	VARCHAR(7)
// )
// use this table to:
// 1. update migration information on upgrade (insert on table)
// 2. update migration information on downgrade (delete from table)
// 3. verify migration current state
//  3.1. if it is consistent (index == count, then OK)
//  3.2. if there is need to upgrade (index < count)
//  3.3. if there is need to downgrade (index > count)
// 4. if it is discontinuous (missing row on database, or some file/checksum diffenrent from loaded file)
//  If that's the case, you can try to recover it by:
//  4.1. Make a backup, just for sure. (data may be lost on the process)
//  4.2. Downgrade until the corrupted point
//   4.2.1. There may be errors, we need our downgrade to ignore them (or just make a function Recover or something)
//  4.3 Upgrade it again
//  4.4 Verify data lost (it may be difficult, I don't think it's possible to do if'ts not manually)

// NOTE: repository package will be refactored soon, there will be some impact on this module.

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/raulscr/arkhanna-master-server/repository"
)

// Private var

const rxUp string = `(?m)(\d+)[-_\.]up[-_\.](?:.+?)\.sql`
const rxDown string = `(?m)(\d+)[-_\.]down[-_\.](?:.+?)\.sql`

var reUp *regexp.Regexp = regexp.MustCompile(rxUp)
var reDown *regexp.Regexp = regexp.MustCompile(rxDown)

// Interfaces

type MigrationInterface interface {
	Upgrade() error
	Downgrade() error
	prepare(migrationPath string) error
}

// Public types

type MigrationError struct {
	error
	FileNameError string
}

type MigratinFilePair struct {
	FileUp   string
	FileDown string
}

type MigrationService struct {
	count          int
	index          int
	migrationPath  string
	migrationFiles map[int]*MigratinFilePair
}

func (e MigrationError) Error() string {
	return fmt.Sprintf("Error on file %s:\n%s", e.FileNameError, e.error.Error())
}

func (m MigrationService) Upgrade() error {
	for m.index < m.count {
		var err = m.runCurrentFileUp()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m MigrationService) Downgrade() error {
	m.index = m.count
	for m.index > 0 {
		var err = m.runCurrentFileDown()
		if err != nil {
			return err
		}
	}

	return nil
}

func NewMigration(migration_path string) (*MigrationService, error) {
	var m *MigrationService = new(MigrationService)
	var err error

	m.migrationFiles = make(map[int]*MigratinFilePair)
	m.migrationPath = migration_path
	m.index = 0
	m.count = 0

	err = m.prepare(migration_path)
	if err != nil {
		m = nil
	}

	return m, err
}

// Private

func (m *MigrationService) prepare(migration_path string) error {

	dir, err := os.Open(migration_path)
	if err != nil {
		return err
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, f := range files {
		if reUp.MatchString(f.Name()) {
			err = m.setFileUp(f.Name())
		} else if reDown.MatchString(f.Name()) {
			err = m.setFileDown(f.Name())
		}

		if err != nil {
			return err
		}
	}

	return m.verify()
}

func (m *MigrationService) verifyIndex(index int) error {
	files, exist := m.migrationFiles[index]
	if !exist || files == nil || files.FileUp == "" || files.FileDown == "" {
		return fmt.Errorf("migration %d not exist", index)
	}
	return nil
}

func (m *MigrationService) verify() error {
	for i := 1; i <= m.count; i++ {
		var err = m.verifyIndex(i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MigrationService) setFileUp(fileName string) error {
	index, err := strconv.ParseInt(reUp.FindAllStringSubmatch(fileName, -1)[0][1], 10, 32) // Verify if the string exists
	if err != nil {
		return err
	}

	_, exists := m.migrationFiles[int(index)]
	if !exists {
		m.migrationFiles[int(index)] = new(MigratinFilePair)
	}

	m.migrationFiles[int(index)].FileUp = fileName
	m.updateFileCount(int(index))
	return nil
}

func (m *MigrationService) setFileDown(fileName string) error {
	index, err := strconv.ParseInt(reDown.FindAllStringSubmatch(fileName, -1)[0][1], 10, 32) // Verify if the string exists
	if err != nil {
		return err
	}

	_, exists := m.migrationFiles[int(index)]
	if !exists {
		m.migrationFiles[int(index)] = new(MigratinFilePair)
	}

	m.migrationFiles[int(index)].FileDown = fileName
	m.updateFileCount(int(index))
	return nil
}

func (m *MigrationService) updateFileCount(index int) {
	if index > m.count {
		m.count = index
	}
}

func (m *MigrationService) runCurrentFileUp() error {
	m.index++
	return m.runFileUp(m.index, true)
}

func (m *MigrationService) runCurrentFileDown() error {
	m.index--
	return m.runFileUp(m.index, false)
}

func (m *MigrationService) runFileUp(index int, isRunUp bool) error {
	var file string
	var err error = m.verifyIndex(index)
	if err != nil {
		return err
	}

	if isRunUp {
		file = m.migrationFiles[index].FileUp
	} else {
		file = m.migrationFiles[index].FileDown
	}

	err = runFile(m.migrationPath + file)

	if err != nil {
		err = MigrationError{error: err, FileNameError: file}
	}

	return err
}

func runFile(file_path string) error {
	file_content, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}

	return repository.ExecRaw(string(file_content))
}

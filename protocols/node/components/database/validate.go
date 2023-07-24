package database

import (
	"fmt"

	iface "github.com/taubyte/go-interfaces/services/substrate/components/database"
)

func (s *Service) validateCommit(hash, projectId, branch string) (bool, string, error) {
	s.commitLock.Lock()
	pastCommit, ok := s.commits[hash]
	s.commitLock.Unlock()
	if !ok {
		return false, "", fmt.Errorf("hash `%s` not found in commit cache", hash)
	}

	newCommit, err := s.Tns().Simple().Commit(projectId, branch)
	if err != nil {
		return false, "", err
	}

	if pastCommit != newCommit {
		return false, "", nil
	}

	return true, newCommit, nil

}

func (s *Service) updateDatabase(database iface.Database) (iface.Database, error) {
	config := database.DBContext()
	newConfig, err := s.Tns().Database().All(config.ProjectId, config.ApplicationId, s.Branch()).GetByName(config.Config.Name)
	if err != nil {
		return nil, err
	}

	database.KV().UpdateSize(newConfig.Size)

	return database, nil
}

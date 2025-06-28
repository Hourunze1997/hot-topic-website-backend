package app

import "errors"

type CmdToAddTopicSolution []OptionalTopic

func (cmd CmdToAddTopicSolution) init() {
	for i := range cmd {
		cmd[i].init()
	}
}

func (cmd CmdToAddTopicSolution) Validate() error {
	for i := range cmd {
		items := cmd[i].DiscussionSources
		for i := range items {
			resolved, unresolved := items[i].filterout()
			if len(unresolved) != 0 && len(resolved) != 1 {
				return errors.New("resolved num is not 1")
			}
		}
	}

	return nil
}

/**
 * Created by super on 2019/3/9.
 *awesomeProject.
 */
package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSER            = "c"
	VIDEO_PATH="./videos/"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error

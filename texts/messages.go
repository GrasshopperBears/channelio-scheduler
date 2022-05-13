package texts

const MESSAGE_ERR_PRFIX string = "오류: "

const MESSAGE_HELP string = "사용할 수 있는 명령어\n" +
	SCHEDULER_PREFIX + " " + SCHEDULER_ADD + " : 현재 그룹에 일정을 추가합니다\n" +
	SCHEDULER_PREFIX + " " + SCHEDULER_GET + " : 현재 그룹에 추가된 일정을 조회합니다\n" +
	SCHEDULER_PREFIX + " " + SCHEDULER_DELETE + " : 해당 일정을 삭제합니다\n"

// Success messeges
const MESSAGE_ADD_SUCCESS = "일정이 추가되었습니다."
const MESSAGE_DELETE_SUCCESS = "일정이 삭제되었습니다."

// Error messeges
const MESSAGE_WRONG_FORMAT string = MESSAGE_ERR_PRFIX + "잘못된 형식입니다. 사용할 수 있는 명령어는 '" + SCHEDULER_PREFIX + " " + SCHEDULER_HELP + "'을 통해 확인할 수 있습니다."
const MESSAGE_NO_SCHEDULE string = MESSAGE_ERR_PRFIX + "일정이 존재하지 않습니다. '" + SCHEDULER_PREFIX + " " + SCHEDULER_ADD + "' 명령어를 통해 일정을 추가해보세요."
const MESSAGE_DELETE_BEFORE_GET string = MESSAGE_ERR_PRFIX + "일정 조회 후 삭제할 수 있습니다."
const MESSAGE_SCHEDULE_NOT_FOUND string = MESSAGE_ERR_PRFIX + "존재하지 않는 일정입니다."

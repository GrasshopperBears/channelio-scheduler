package texts

const MESSAGE_WRONG_FORMAT string = "잘못된 형식입니다. 사용할 수 있는 명령어는 '" + SCHEDULER_PREFIX + " " + SCHEDULER_HELP + "'을 통해 확인할 수 있습니다."
const MESSAGE_HELP string = "사용할 수 있는 명령어\n" +
	SCHEDULER_PREFIX + " " + SCHEDULER_ADD + " : 현재 그룹에 일정을 추가합니다\n" +
	SCHEDULER_PREFIX + " " + SCHEDULER_GET + " : 현재 그룹에 추가된 일정을 조회합니다\n" +
	SCHEDULER_PREFIX + " " + SCHEDULER_DELETE + " : 해당 일정을 삭제합니다\n"
const MESSAGE_NO_SCHEDULE string = "일정이 존재하지 않습니다. '" + SCHEDULER_PREFIX + " " + SCHEDULER_ADD + "' 명령어를 통해 일정을 추가해보세요."

package translation

const TaskQueueName = "translation-tasks"

type TranslationWorkflowInput struct {
	Name         string
	LanguageCode string
}

type TranslationWorkflowOutput struct {
	HelloMessage   string
	GoodbyeMessage string
}

// TODO define structs for Activity input and output here
type TranslationActivityInput struct {
	Term         string
	LanguageCode string
}

type TranslationActivityOutput struct {
	Translated string
}

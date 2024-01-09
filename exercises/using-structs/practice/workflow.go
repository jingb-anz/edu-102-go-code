package translation

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloGoodbye(ctx workflow.Context, input TranslationWorkflowInput) (TranslationWorkflowOutput, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 45,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// TODO Create your Activity input struct and populate it with the last
	//      two fields from the ExecuteActivity call below
	helloActivityInput := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: input.LanguageCode,
	}

	// TODO Replace "string" below with your Activity output struct type
	var helloActivityResult TranslationActivityOutput

	// TODO Use your input struct in the ExecuteActivity call below
	err := workflow.ExecuteActivity(ctx, TranslateTerm, helloActivityInput).Get(ctx, &helloActivityResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	// TODO Update the middle parameter to use the Translation field from the Activity output struct
	helloMessage := fmt.Sprintf("%s, %s", helloActivityResult, input.Name)

	// TODO Create your Activity input struct and populate it with the last
	//      two fields from the ExecuteActivity call below
	goodbyeActivityInput := TranslationActivityInput{
		Term:         "GoodBye",
		LanguageCode: input.LanguageCode,
	}

	// TODO Replace "string" below with your Activity output struct type
	var goodbyeActivityResult TranslationActivityOutput

	// TODO Use your input struct in the ExecuteActivity call below
	err = workflow.ExecuteActivity(ctx, TranslateTerm, goodbyeActivityInput).Get(ctx, &goodbyeActivityResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	// TODO Update the middle parameter to use the Translation field from the Activity output struct
	goodbyeMessage := fmt.Sprintf("%s, %s", goodbyeActivityResult, input.Name)

	output := TranslationWorkflowOutput{
		HelloMessage:   helloMessage,
		GoodbyeMessage: goodbyeMessage,
	}

	return output, nil
}

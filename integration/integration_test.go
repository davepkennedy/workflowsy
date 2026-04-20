//go:build feature

package integration

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"workflowsy/internal"

	"github.com/cucumber/godog"
)

var (
	ErrNoCalculator = errors.New("no calculator found")
	ErrNoInput	  = errors.New("no input found")
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "calculator",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features"},
			TestingT: t,
			Strict: true,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Given(
		`^I have a calculator$`,
		func(ctx context.Context) context.Context {
			calculator := internal.NewCalculator()
			return context.WithValue(ctx, "calculator", calculator)
		},
	)
	sc.When(
		`^I input "(.+)"$`,
		func(ctx context.Context, input string) context.Context {
			return context.WithValue(ctx, "input", input)
		},	
	)
	sc.Then(
		`^the result should be (\-?\d+.?\d*)$`,
		func(ctx context.Context, expected float64) error {
			calculator, ok := ctx.Value("calculator").(internal.Calculator)
			if !ok {
				return ErrNoCalculator
			}

			input, ok := ctx.Value("input").(string)
			if !ok {
				return ErrNoInput
			}

			val, err := calculator.Process(input)
			if err != nil {
				return err
			}
			if val != expected {
				return fmt.Errorf("expected %f but got %f", expected, val)
			}
			return nil
		},
	)
}
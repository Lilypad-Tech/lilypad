package system

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

type CommandContext struct {
	CommandContext context.Context
	Ctx            context.Context
	Cm             *CleanupManager
	CancelFunc     context.CancelFunc
}

func NewSystemContext(ctx context.Context) *CommandContext {
	SetupLogging()
	cm := NewCleanupManager()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	return &CommandContext{
		CommandContext: ctx,
		Ctx:            ctx,
		Cm:             cm,
		CancelFunc:     cancel,
	}
}

func NewTestingContext() *CommandContext {
	go func() {
		// Create a new channel to receive signal notifications.
		sigChan := make(chan os.Signal, 1)

		// Notify sends the incoming signals to the sigChan.
		// We're only interested in SIGINT (CTRL+C) here,
		// but you can add others if needed.
		signal.Notify(sigChan, syscall.SIGINT)

		// Wait for a signal to be caught.
		sig := <-sigChan

		// Handle the received signal.
		fmt.Printf("Caught signal %s: shutting down.\n", sig)

		// Perform your cleanup logic here,
		// like closing database connections, terminating network activities, etc.

		// Finally, exit the program.
		os.Exit(0)
	}()
	return NewSystemContext(context.Background())
}

func NewCommandContext(cmd *cobra.Command) *CommandContext {
	return NewSystemContext(cmd.Context())
}

func (cmdContext *CommandContext) Cleanup() {
	cmdContext.Cm.Cleanup(cmdContext.CommandContext)
	cmdContext.Cm.Cleanup(cmdContext.Ctx)
	cmdContext.CancelFunc()
}

// NewDetachedContext produces a new context that has a separate cancellation mechanism from its parent. This should be
// used when having to continue using a context after it has been canceled, such as cleaning up Docker resources
// after the context has been canceled but want to keep work associated with the same trace.
func NewDetachedContext(parent context.Context) context.Context {
	return detachedContext{parent: parent}
}

var _ context.Context = detachedContext{}

type detachedContext struct {
	parent context.Context
}

func (d detachedContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (d detachedContext) Done() <-chan struct{} {
	return nil
}

func (d detachedContext) Err() error {
	return nil
}

func (d detachedContext) Value(key any) any {
	return d.parent.Value(key)
}

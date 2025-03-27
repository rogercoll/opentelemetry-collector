package pipeline

import (
	"context"
	"time"
)

type pipelineStartContextKey struct{}

func RecordPipelineStart(ctx context.Context) context.Context {
	return context.WithValue(ctx, pipelineStartContextKey{}, []time.Time{time.Now()})
}

func PipelineDuration(ctx context.Context) []int64 {
	if startTimes, ok := ctx.Value(pipelineStartContextKey{}).([]time.Time); ok {
		pipelineTimes := make([]int64, len(startTimes))
		for i := range startTimes {
			pipelineTimes[i] = time.Since(startTimes[i]).Milliseconds()
		}
		return pipelineTimes
	}
	return nil
}

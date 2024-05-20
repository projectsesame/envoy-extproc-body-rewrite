package main

import (
	ep "github.com/wrossmorrow/envoy-extproc-sdk-go"
	"log"
)

type bodyRewriteRequestProcessor struct {
	opts        *ep.ProcessingOptions
	bodyRewrite string
}

func (s *bodyRewriteRequestProcessor) GetName() string {
	return "body-rewrite"
}

func (s *bodyRewriteRequestProcessor) GetOptions() *ep.ProcessingOptions {
	return s.opts
}

func (s *bodyRewriteRequestProcessor) ProcessRequestHeaders(ctx *ep.RequestContext, headers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *bodyRewriteRequestProcessor) ProcessRequestBody(ctx *ep.RequestContext, body []byte) error {
	ctx.ReplaceBodyChunk([]byte(s.bodyRewrite))
	return ctx.ContinueRequest()
}

func (s *bodyRewriteRequestProcessor) ProcessRequestTrailers(ctx *ep.RequestContext, trailers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *bodyRewriteRequestProcessor) ProcessResponseHeaders(ctx *ep.RequestContext, headers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

func (s *bodyRewriteRequestProcessor) ProcessResponseBody(ctx *ep.RequestContext, body []byte) error {
	return ctx.ContinueRequest()
}

func (s *bodyRewriteRequestProcessor) ProcessResponseTrailers(ctx *ep.RequestContext, trailers ep.AllHeaders) error {
	return ctx.ContinueRequest()
}

const kbodyRewrite = "body-rewrite"

func (s *bodyRewriteRequestProcessor) Init(opts *ep.ProcessingOptions, nonFlagArgs []string) error {
	s.opts = opts
	s.bodyRewrite = "this is a replacement body."

	var (
		i           int
		bodyRewrite string
	)

	nArgs := len(nonFlagArgs)
	for ; i < nArgs-1; i++ {
		if nonFlagArgs[i] == kbodyRewrite {
			break
		}
	}

	if i == nArgs {
		log.Printf("the argument: 'body-rewrite' is missing, use the default.\n")
		return nil
	}

	bodyRewrite = nonFlagArgs[i+1]

	if len(bodyRewrite) > 0 {
		s.bodyRewrite = bodyRewrite
		log.Printf("the body rewrite is: %s\n", s.bodyRewrite)
	}

	return nil
}

func (s *bodyRewriteRequestProcessor) Finish() {}

package clock

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FromHeader(h http.Header) (vc VClock, err error) {
	header, ok := h["V-Clock"]
	if !ok {
		return nil, ErrHeaderMissing
	}

	vcHeaders := make([]string, 0)
	for _, value := range header {
		for _, entry := range strings.Split(value, ", ") {
			vcHeaders = append(vcHeaders, entry)
		}
	}

	vc = New()
	for _, value := range vcHeaders {
		data := strings.SplitN(value, "=", 2)
		if len(data) != 2 {
			return nil, ErrHeaderParseFormat
		}

		nodeID, err := uuid.Parse(data[0])
		if err != nil {
			return nil, ErrHeaderParseNodeID
		}

		tick, err := strconv.ParseUint(data[1], 10, 64)
		if err != nil {
			return nil, ErrHeaderParseTick
		}

		vc.SetTick(nodeID, tick)
	}

	return vc, nil
}

func (v VClock) IntoHeader() http.Header {
	header := http.Header{}
	for nodeID, tick := range v {
		header.Add("V-Clock", fmt.Sprintf("%s=%d", nodeID.String(), tick))
	}
	return header
}

var (
	ErrHeaderMissing     = errors.New("missing \"V-Clock\" header")
	ErrHeaderParseFormat = errors.New("error parsing \"V-Clock\" header format")
	ErrHeaderParseNodeID = errors.New("error parsing \"V-Clock\" header node ID")
	ErrHeaderParseTick   = errors.New("error parsing \"V-Clock\" header tick")
)

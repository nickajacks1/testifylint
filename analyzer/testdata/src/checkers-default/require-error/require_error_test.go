// Code generated by testifylint/internal/testgen. DO NOT EDIT.

package requireerror

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestRequireErrorChecker(t *testing.T) {
	var err error

	assObj, reqObj := assert.New(t), require.New(t)

	// Invalid.
	assert.Error(t, err)                                                          // want "require-error: for error assertions use the `require`"
	assert.Errorf(t, err, "msg with args %d %s", 42, "42")                        // want "require-error: for error assertions use the `require`"
	assert.ErrorIs(t, err, io.EOF)                                                // want "require-error: for error assertions use the `require`"
	assert.ErrorIsf(t, err, io.EOF, "msg with args %d %s", 42, "42")              // want "require-error: for error assertions use the `require`"
	assert.ErrorAs(t, err, new(os.PathError))                                     // want "require-error: for error assertions use the `require`"
	assert.ErrorAsf(t, err, new(os.PathError), "msg with args %d %s", 42, "42")   // want "require-error: for error assertions use the `require`"
	assert.EqualError(t, err, "end of file")                                      // want "require-error: for error assertions use the `require`"
	assert.EqualErrorf(t, err, "end of file", "msg with args %d %s", 42, "42")    // want "require-error: for error assertions use the `require`"
	assert.ErrorContains(t, err, "end of file")                                   // want "require-error: for error assertions use the `require`"
	assert.ErrorContainsf(t, err, "end of file", "msg with args %d %s", 42, "42") // want "require-error: for error assertions use the `require`"
	assert.NoError(t, err)                                                        // want "require-error: for error assertions use the `require`"
	assert.NoErrorf(t, err, "msg with args %d %s", 42, "42")                      // want "require-error: for error assertions use the `require`"
	assert.NotErrorIs(t, err, io.EOF)                                             // want "require-error: for error assertions use the `require`"
	assert.NotErrorIsf(t, err, io.EOF, "msg with args %d %s", 42, "42")           // want "require-error: for error assertions use the `require`"
	assObj.Error(err)                                                             // want "require-error: for error assertions use the `require`"
	assObj.Errorf(err, "msg with args %d %s", 42, "42")                           // want "require-error: for error assertions use the `require`"
	assObj.ErrorIs(err, io.EOF)                                                   // want "require-error: for error assertions use the `require`"
	assObj.ErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")                 // want "require-error: for error assertions use the `require`"
	assObj.ErrorAs(err, new(os.PathError))                                        // want "require-error: for error assertions use the `require`"
	assObj.ErrorAsf(err, new(os.PathError), "msg with args %d %s", 42, "42")      // want "require-error: for error assertions use the `require`"
	assObj.EqualError(err, "end of file")                                         // want "require-error: for error assertions use the `require`"
	assObj.EqualErrorf(err, "end of file", "msg with args %d %s", 42, "42")       // want "require-error: for error assertions use the `require`"
	assObj.ErrorContains(err, "end of file")                                      // want "require-error: for error assertions use the `require`"
	assObj.ErrorContainsf(err, "end of file", "msg with args %d %s", 42, "42")    // want "require-error: for error assertions use the `require`"
	assObj.NoError(err)                                                           // want "require-error: for error assertions use the `require`"
	assObj.NoErrorf(err, "msg with args %d %s", 42, "42")                         // want "require-error: for error assertions use the `require`"
	assObj.NotErrorIs(err, io.EOF)                                                // want "require-error: for error assertions use the `require`"
	assObj.NotErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")              // want "require-error: for error assertions use the `require`"

	// Valid.
	require.Error(t, err)
	require.Errorf(t, err, "msg with args %d %s", 42, "42")
	require.ErrorIs(t, err, io.EOF)
	require.ErrorIsf(t, err, io.EOF, "msg with args %d %s", 42, "42")
	require.ErrorAs(t, err, new(os.PathError))
	require.ErrorAsf(t, err, new(os.PathError), "msg with args %d %s", 42, "42")
	require.EqualError(t, err, "end of file")
	require.EqualErrorf(t, err, "end of file", "msg with args %d %s", 42, "42")
	require.ErrorContains(t, err, "end of file")
	require.ErrorContainsf(t, err, "end of file", "msg with args %d %s", 42, "42")
	require.NoError(t, err)
	require.NoErrorf(t, err, "msg with args %d %s", 42, "42")
	require.NotErrorIs(t, err, io.EOF)
	require.NotErrorIsf(t, err, io.EOF, "msg with args %d %s", 42, "42")
	reqObj.Error(err)
	reqObj.Errorf(err, "msg with args %d %s", 42, "42")
	reqObj.ErrorIs(err, io.EOF)
	reqObj.ErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")
	reqObj.ErrorAs(err, new(os.PathError))
	reqObj.ErrorAsf(err, new(os.PathError), "msg with args %d %s", 42, "42")
	reqObj.EqualError(err, "end of file")
	reqObj.EqualErrorf(err, "end of file", "msg with args %d %s", 42, "42")
	reqObj.ErrorContains(err, "end of file")
	reqObj.ErrorContainsf(err, "end of file", "msg with args %d %s", 42, "42")
	reqObj.NoError(err)
	reqObj.NoErrorf(err, "msg with args %d %s", 42, "42")
	reqObj.NotErrorIs(err, io.EOF)
	reqObj.NotErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")
}

type RequireErrorCheckerSuite struct {
	suite.Suite
}

func TestRequireErrorCheckerSuite(t *testing.T) {
	suite.Run(t, new(RequireErrorCheckerSuite))
}

func (s *RequireErrorCheckerSuite) TestAll() {
	var err error

	assObj, reqObj := s.Assert(), s.Require()

	// Invalid.
	s.Error(err)                                                                   // want "require-error: for error assertions use the `require`"
	s.Errorf(err, "msg with args %d %s", 42, "42")                                 // want "require-error: for error assertions use the `require`"
	s.ErrorIs(err, io.EOF)                                                         // want "require-error: for error assertions use the `require`"
	s.ErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")                       // want "require-error: for error assertions use the `require`"
	s.ErrorAs(err, new(os.PathError))                                              // want "require-error: for error assertions use the `require`"
	s.ErrorAsf(err, new(os.PathError), "msg with args %d %s", 42, "42")            // want "require-error: for error assertions use the `require`"
	s.EqualError(err, "end of file")                                               // want "require-error: for error assertions use the `require`"
	s.EqualErrorf(err, "end of file", "msg with args %d %s", 42, "42")             // want "require-error: for error assertions use the `require`"
	s.ErrorContains(err, "end of file")                                            // want "require-error: for error assertions use the `require`"
	s.ErrorContainsf(err, "end of file", "msg with args %d %s", 42, "42")          // want "require-error: for error assertions use the `require`"
	s.NoError(err)                                                                 // want "require-error: for error assertions use the `require`"
	s.NoErrorf(err, "msg with args %d %s", 42, "42")                               // want "require-error: for error assertions use the `require`"
	s.NotErrorIs(err, io.EOF)                                                      // want "require-error: for error assertions use the `require`"
	s.NotErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")                    // want "require-error: for error assertions use the `require`"
	s.Assert().Error(err)                                                          // want "require-error: for error assertions use the `require`"
	s.Assert().Errorf(err, "msg with args %d %s", 42, "42")                        // want "require-error: for error assertions use the `require`"
	s.Assert().ErrorIs(err, io.EOF)                                                // want "require-error: for error assertions use the `require`"
	s.Assert().ErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")              // want "require-error: for error assertions use the `require`"
	s.Assert().ErrorAs(err, new(os.PathError))                                     // want "require-error: for error assertions use the `require`"
	s.Assert().ErrorAsf(err, new(os.PathError), "msg with args %d %s", 42, "42")   // want "require-error: for error assertions use the `require`"
	s.Assert().EqualError(err, "end of file")                                      // want "require-error: for error assertions use the `require`"
	s.Assert().EqualErrorf(err, "end of file", "msg with args %d %s", 42, "42")    // want "require-error: for error assertions use the `require`"
	s.Assert().ErrorContains(err, "end of file")                                   // want "require-error: for error assertions use the `require`"
	s.Assert().ErrorContainsf(err, "end of file", "msg with args %d %s", 42, "42") // want "require-error: for error assertions use the `require`"
	s.Assert().NoError(err)                                                        // want "require-error: for error assertions use the `require`"
	s.Assert().NoErrorf(err, "msg with args %d %s", 42, "42")                      // want "require-error: for error assertions use the `require`"
	s.Assert().NotErrorIs(err, io.EOF)                                             // want "require-error: for error assertions use the `require`"
	s.Assert().NotErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")           // want "require-error: for error assertions use the `require`"
	assObj.Error(err)                                                              // want "require-error: for error assertions use the `require`"
	assObj.Errorf(err, "msg with args %d %s", 42, "42")                            // want "require-error: for error assertions use the `require`"
	assObj.ErrorIs(err, io.EOF)                                                    // want "require-error: for error assertions use the `require`"
	assObj.ErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")                  // want "require-error: for error assertions use the `require`"
	assObj.ErrorAs(err, new(os.PathError))                                         // want "require-error: for error assertions use the `require`"
	assObj.ErrorAsf(err, new(os.PathError), "msg with args %d %s", 42, "42")       // want "require-error: for error assertions use the `require`"
	assObj.EqualError(err, "end of file")                                          // want "require-error: for error assertions use the `require`"
	assObj.EqualErrorf(err, "end of file", "msg with args %d %s", 42, "42")        // want "require-error: for error assertions use the `require`"
	assObj.ErrorContains(err, "end of file")                                       // want "require-error: for error assertions use the `require`"
	assObj.ErrorContainsf(err, "end of file", "msg with args %d %s", 42, "42")     // want "require-error: for error assertions use the `require`"
	assObj.NoError(err)                                                            // want "require-error: for error assertions use the `require`"
	assObj.NoErrorf(err, "msg with args %d %s", 42, "42")                          // want "require-error: for error assertions use the `require`"
	assObj.NotErrorIs(err, io.EOF)                                                 // want "require-error: for error assertions use the `require`"
	assObj.NotErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")               // want "require-error: for error assertions use the `require`"

	// Valid.
	s.Require().Error(err)
	s.Require().Errorf(err, "msg with args %d %s", 42, "42")
	s.Require().ErrorIs(err, io.EOF)
	s.Require().ErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")
	s.Require().ErrorAs(err, new(os.PathError))
	s.Require().ErrorAsf(err, new(os.PathError), "msg with args %d %s", 42, "42")
	s.Require().EqualError(err, "end of file")
	s.Require().EqualErrorf(err, "end of file", "msg with args %d %s", 42, "42")
	s.Require().ErrorContains(err, "end of file")
	s.Require().ErrorContainsf(err, "end of file", "msg with args %d %s", 42, "42")
	s.Require().NoError(err)
	s.Require().NoErrorf(err, "msg with args %d %s", 42, "42")
	s.Require().NotErrorIs(err, io.EOF)
	s.Require().NotErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")
	reqObj.Error(err)
	reqObj.Errorf(err, "msg with args %d %s", 42, "42")
	reqObj.ErrorIs(err, io.EOF)
	reqObj.ErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")
	reqObj.ErrorAs(err, new(os.PathError))
	reqObj.ErrorAsf(err, new(os.PathError), "msg with args %d %s", 42, "42")
	reqObj.EqualError(err, "end of file")
	reqObj.EqualErrorf(err, "end of file", "msg with args %d %s", 42, "42")
	reqObj.ErrorContains(err, "end of file")
	reqObj.ErrorContainsf(err, "end of file", "msg with args %d %s", 42, "42")
	reqObj.NoError(err)
	reqObj.NoErrorf(err, "msg with args %d %s", 42, "42")
	reqObj.NotErrorIs(err, io.EOF)
	reqObj.NotErrorIsf(err, io.EOF, "msg with args %d %s", 42, "42")
}

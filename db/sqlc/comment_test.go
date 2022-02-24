package db

import (
	"context"
	"testing"

	"github.com/harissucipto/xendit-task/util"
	"github.com/stretchr/testify/require"
)


func createRandomComment(t *testing.T, orgName string) Comment {

	arg := CreateCommentParams{
		Content: util.RandomComment(),
		OrgName: orgName,
	}
	comment, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, comment)

	return comment
}

func TestCreateComment(t *testing.T) {
	for i := 0; i < 5; i++ {
		orgName := util.RandomString(6)
		comment := createRandomComment(t, orgName)
		require.NotNil(t, comment)
	}
}

func TestDeleteComment(t *testing.T) {
	orgName := util.RandomString(6)
	comment := createRandomComment(t, orgName)
	require.NotNil(t, comment)

	err := testQueries.DeleteComment(context.Background(), orgName)
	require.NoError(t, err)
}

func TestListComments(t *testing.T) {
	orgName := util.RandomString(6)
	for i := 0; i < 5; i++ {
		comment := createRandomComment(t, orgName)
		require.NotNil(t, comment)
	}

	comments, err := testQueries.ListComments(context.Background(), orgName)
	require.NotEmpty(t, comments)
	require.NoError(t, err)
	require.Len(t, comments, 5)

}

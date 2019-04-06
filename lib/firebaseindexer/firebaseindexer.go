package firebaseindexer

import (
	"context"
	"os"
	"strconv"

	firebase "firebase.google.com/go"
	"github.com/leomfelicissimo/biblionify/lib/types"
	"github.com/leomfelicissimo/biblionify/lib/util"
	"google.golang.org/api/option"
)

// IndexBiblionData get a instance of firebaseApp, biblionTexts json array
// and index the data on firestore at firebase.com
func IndexBiblionData(biblionTexts []types.BiblionText) {
	ctx := context.Background()
	credentialsFilePath := os.Getenv("CREDENTIALS_FILE_PATH")
	opt := option.WithCredentialsFile(credentialsFilePath)
	app, err := firebase.NewApp(ctx, nil, opt)
	util.HandleError(err, "Error indexing biblionText at Firebase Firestore")

	firestore, err := app.Firestore(ctx)
	util.HandleError(err, "Error initializing firestore")

	cRef := firestore.Collection("biblionTexts")

	for _, text := range biblionTexts {
		id := text.Book + strconv.Itoa(text.Chapter) + strconv.Itoa(text.Verse)
		_, err := cRef.Doc(id).Set(ctx, text)
		util.HandleError(err, "Error adding doc to firebase")
	}
}

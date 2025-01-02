package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestWatchAuction(t *testing.T) {
	mongoT := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	auction := auction_entity.Auction{
		Id:          "1",
		ProductName: "Product 1",
		Category:    "Category 1",
		Description: "Description 1",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	durationInterval := time.Duration(time.Second)

	mongoT.Run("test watch auction", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		repo := NewAuctionRepository(mt.DB)
		repo.auctionInterval = durationInterval
		fmt.Printf("repo.auctionInterval: %v\n", repo.auctionInterval)

		go repo.watchAuction(context.Background(), &auction)
		startedEvents := mt.GetAllStartedEvents()
		if len(startedEvents) != 0 {
			mt.Error("expected no events to be started before auction interval")
		}
		time.Sleep(durationInterval + 30*time.Millisecond)
		startedEvents = mt.GetAllStartedEvents()
		if len(startedEvents) == 0 {
			mt.Error("expected events to be started after auction interval")
		}
		array, ok := mt.GetStartedEvent().Command.Lookup("updates").ArrayOK()
		if !ok {
			mt.Fatal("expected updates to be an array")
		}
		firstUpdate, err := array.IndexErr(0)
		if err != nil {
			mt.Fatalf("expected array to have at least one element: %v", err)
		}
		updateDoc, ok := firstUpdate.Value().Document().Lookup("u").Document().Lookup("$set").DocumentOK()
		if !ok {
			mt.Fatal("expected $set to be a document")
		}
		capturedStatus := updateDoc.Lookup("status").Int32()
		if auction_entity.AuctionStatus(capturedStatus) != auction_entity.Completed {
			mt.Errorf("expected status to be %v, got %v", auction_entity.Completed, auction_entity.AuctionStatus(capturedStatus))
		}
	})
}

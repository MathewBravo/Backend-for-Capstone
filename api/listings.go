package endpoints

import (
	"Backend_for_Capstone/api/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Listing struct {
	ListingId     int64   `json:"user_id"`
	ListedBy      string  `json:"listed_by"`
	Price         float64 `json:"price"`
	Rooms         int32   `json:"rooms"`
	Bathrooms     int32   `json:"bathrooms"`
	Description   string  `json:"description"`
	PropertyType  string  `json:"property_type"`
	BuildingType  string  `json:"building_type"`
	Stories       string  `json:"stories"`
	CommunityName string  `json:"community_name"`
	Landsize      string  `json:"landsize"`
	PropertyTax   float64 `json:"property_tax"`
	ParkingType   string  `json:"parking_type"`
	HeatingType   string  `json:"heating_type"`
	CoolingType   string  `json:"cooling_type"`
	AmenitiesNear string  `json:"amenities_near"`
	Street        string  `json:"street"`
	City          string  `json:"city"`
	County        string  `json:"county"`
	PostCode      string  `json:"post_code"`
	Country       string  `json:"country"`
	ReviewStatus  string  `json:"review_status"`
}

type ListingForAllSearch struct {
	ListingId    int64   `json:"user_id"`
	ListedBy     string  `json:"listed_by"`
	Price        float64 `json:"price"`
	PropertyType string  `json:"property_type"`
	Stories      string  `json:"stories"`
}

type ListingsForAdminPending struct {
	ListingId    int64  `json:"listing_id"`
	Street       string `json:"street"`
	City         string `json:"city"`
	ReviewStatus string `json:"review_status"`
	PostedBy     string `json:"posted_by"`
}

func GetMessage(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func GetListingsForAdminPending(c *gin.Context) {
	db := database.Database
	var listings []*ListingsForAdminPending
	query, e := db.Query("select listings.listingid, listings.street, listings.city, listings.review_status, users.username from listings join users on listings.listedby = users.userid where listings.review_status = 'pending';")
	if e != nil {
		log.Fatalf("Database not found")
	}
	for query.Next() {
		l := new(ListingsForAdminPending)
		if err := query.Scan(&l.ListingId, &l.Street, &l.City, &l.ReviewStatus, &l.PostedBy); err != nil {
			log.Fatalf("Errow: %s", err)
		}
		listings = append(listings, l)
	}
	c.JSON(http.StatusOK, listings)

}

func GetListingsForSearchDisplay(c *gin.Context) {
	db := database.Database
	var listings []*ListingForAllSearch
	query, e := db.Query("select listingid, listedby, price, propertytype, storeys  from listings")
	if e != nil {
		log.Fatalf("Database not found")
	}
	for query.Next() {

		l := new(ListingForAllSearch)
		if err := query.Scan(&l.ListingId, &l.ListedBy, &l.Price, &l.PropertyType, &l.Stories); err != nil {
			log.Fatalf("Errow: %s", err)
		}
		listings = append(listings, l)
	}
	c.JSON(http.StatusOK, listings)
}

func GetListingByParam(c *gin.Context) {
	//db := database.Database
	sentParams := c.Request.URL.Query()
	if value, ok := sentParams["id"]; ok {
		log.Println(value)
		getListingsById(c)
	} else {
		fmt.Println("Not Included")
	}
}

func getListingsById(c *gin.Context) {

}

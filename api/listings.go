package endpoints

import (
	"Backend_for_Capstone/api/database"
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
}

type ListingForAllSearch struct {
	ListingId    int64   `json:"user_id"`
	ListedBy     string  `json:"listed_by"`
	Price        float64 `json:"price"`
	PropertyType string  `json:"property_type"`
	Stories      string  `json:"stories"`
}

func GetMessage(c *gin.Context) {
	c.String(http.StatusOK, "hello")
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

	c.JSON(http.StatusOK, listings[0])
}

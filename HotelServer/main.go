package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	pb "final_project_backend/pbGenerated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addrHotel = flag.String("addrHotel", "localhost:8080", "the address to connect to")
)

func login(c *gin.Context) {
	flag.Parse()
	var req pb.LoginUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body request",
		})
		return
	}

	conn, err := grpc.Dial(*addrHotel, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewUsersServiceClient(conn)

	resp, err := client.LoginUser(c, &req)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response = pb.LoginUserResponse{
		Message: resp.Message,
		Success: resp.Success,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func signUp(c *gin.Context) {
	flag.Parse()
	var req pb.SignUpUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body request",
		})
		return
	}

	conn, err := grpc.Dial(*addrHotel, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewUsersServiceClient(conn)

	resp, err := client.SignUpUser(c, &req)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response = pb.SignUpUserResponse{
		Message: resp.Message,
		Success: resp.Success,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func addCredit(c *gin.Context) {
	flag.Parse()
	var req pb.AddCreditRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body request",
		})
		return
	}

	conn, err := grpc.Dial(*addrHotel, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewUsersServiceClient(conn)

	resp, err := client.AddCredit(c, &req)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response = pb.AddCreditResponse{
		Message: resp.Message,
		Credit:  resp.Credit,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func unavailableDates(c *gin.Context) {
	flag.Parse()
	var req pb.UnavailableDatesRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body request",
		})
		return
	}

	conn, err := grpc.Dial(*addrHotel, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewUsersServiceClient(conn)

	resp, err := client.UnavailableDates(c, &req)
	if err != nil {
		log.Printf("could not process request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var response = pb.UnavailableDatesResponse{
		Dates: resp.Dates,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.POST("/sign-up", signUp)
	router.POST("/login", login)
	router.POST("/add-credit", addCredit)
	router.POST("/unavailable-dates", unavailableDates)
	router.Run("localhost:8000")
}

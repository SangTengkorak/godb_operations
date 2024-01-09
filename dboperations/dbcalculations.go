package dboperations

import (
	"database/sql"
	"fmt"
	"log"

	"mastengkorak.com/godatabasebasic/datatypes"
)

func Mostexpensive() {
	mahal := DB.QueryRow("SELECT product_id, product_name, product_price FROM stored_products ORDER BY product_price DESC LIMIT 1")
	var mhl datatypes.StoredProduct
	switch err := mahal.Scan(&mhl.Product_id, &mhl.Product_name, &mhl.Product_price); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		fmt.Printf("Most Expensive item is ID : %d, Name: %s, Price: %.2f\n", mhl.Product_id, mhl.Product_name, mhl.Product_price)
	default:
		panic(err)
	}
}

func Cheapest() {
	murah := DB.QueryRow("SELECT product_id, product_name, product_price FROM stored_products ORDER BY product_price ASC LIMIT 1")
	var mrh datatypes.StoredProduct
	switch err := murah.Scan(&mrh.Product_id, &mrh.Product_name, &mrh.Product_price); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		fmt.Printf("Cheapest item is ID : %d, Name: %s, Price: %.2f\n", mrh.Product_id, mrh.Product_name, mrh.Product_price)
	default:
		panic(err)
	}
}

func Allitems() {
	alrows, err := DB.Query("SELECT product_id, product_name, product_price, purchased_date, in_stock FROM stored_products")
	if err != nil {
		log.Fatal(err)
	}

	var strdproducts []datatypes.StoredProduct
	for alrows.Next() {
		var strdprod datatypes.StoredProduct
		err := alrows.Scan(&strdprod.Product_id, &strdprod.Product_name, &strdprod.Product_price, &strdprod.Purchased_date, &strdprod.In_stock)
		if err != nil {
			log.Fatal(err)
		}
		strdproducts = append(strdproducts, strdprod)
	}

	if err := alrows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nStored Products:")
	for _, prod := range strdproducts {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f, Purchased Date: %s, In Stock: %d\n", prod.Product_id, prod.Product_name, prod.Product_price, prod.Purchased_date, prod.In_stock)
	}
	defer alrows.Close()

}

func Buyitem(typedPID int, savings float64) {
	beli := DB.QueryRow("SELECT product_id, product_name, product_price FROM stored_products WHERE product_id = ?", typedPID)
	var bli datatypes.StoredProduct
	switch err := beli.Scan(&bli.Product_id, &bli.Product_name, &bli.Product_price); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		fmt.Printf("\nFound Item with ID of : %d, with Name: %s, and Price: %.2f\n", bli.Product_id, bli.Product_name, bli.Product_price)
	default:
		panic(err)
	}

	if savings < bli.Product_price {
		fmt.Println("\nYour savings not enough to buy this product.")
	} else if savings >= bli.Product_price {
		savings_return := savings - bli.Product_price
		fmt.Printf("You will receive Rp. %.2f from your purchase.", savings_return)
	}
}

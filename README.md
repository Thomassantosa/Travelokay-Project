# Travelokay-Project

| Endpoint	    | Access	| Method	| Function          |
| ------------- | --------- | --------- |------------------ |
| /login	    | All		| POST		| Login             |
| /logout	    | All		| GET		| Logout            |
| /user		    | User		| POST		| AddNewUser	    |
| /user		    | User		| PUT		| UpdateUser	    |
| /user/hotel	| User		| GET		| GetHotelList      |
| /user/hotel	| User		| POST		| AddNewHotelOrder  |
| /user/flight	| User		| GET		| GetFlightList     |
| /user/flight	| User		| POST		| AddNewFlightOrder |
| /user/bus 	| User		| GET		| GetBusList        |
| /user/bus 	| User		| POST		| AddNewBusOrder    |
| /user/train	| User		| GET		| GetTrainList      |
| /user/train	| User		| POST		| AddNewTrainOrder  |
| /user/tour	| User		| GET		| GetTourList       |
| /user/tour	| User		| POST		| AddNewTourOrder   |
| /user/order	| User		| GET		| GetUserOrder      |
| /user/order	| User		| DELETE	| RequestRefund     |
| ------------- | --------- | --------- |------------------ |
| /mitra	    | Mitra		| POST		| AddNewMitra	    |
| /mitra	    | Mitra		| PUT		| UpdateMitra   	|
| /mitra/product| Mitra		| GET		| GetProductList   	|
| /mitra/product| Mitra		| POST  	| AddNewProduct   	|
| /mitra/product| Mitra		| DELETE	| DeleteProduct   	|
| /admin/refund | Mitra		| GET		| GetRefundList   	|
| /admin/refund | Mitra		| DELETE	| ApproveRefund   	|
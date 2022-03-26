# Travelokay-Project

| No | Endpoint	        | Access	| Method	| Function          | Response          | Desc                                                          |
|----| ---------------- | --------- | --------- | ----------------- | ----------------- | ------------------------------------------------------------- |
| 1  | /login	        | All		| POST		| Login             | MessageResponse   | Saat login berhasil, akan menghasilkan cookie berisi token.   |
| 2  | /logout	        | All		| GET		| Logout            | MessageResponse   | Menghapus cookie yang berisi token.                           |
| 3  | /user		    | User		| POST		| AddNewUser	    | MessageResponse   | Register user.                                                |
| 4  | /user		    | User		| PUT		| UpdateUser	    | MessageResponse   | -                                                             |
| 5  | /user/hotel	    | User		| GET		| GetHotelList      | HotelsResponse    | -                                                             |
| 6  | /user/hotel	    | User		| POST		| AddNewHotelOrder  | MessageResponse   | Trigger table rooms.                                          |
| 7  | /user/flight	    | User		| GET		| GetFlightList     | FlightResponse    | -                                                             |
| 8  | /user/flight	    | User		| POST		| AddNewFlightOrder | MessageResponse   | Trigger table seats.                                          |
| 9  | /user/bus 	    | User		| GET		| GetBusList        | BussesResponse    | -                                                             |
| 10 | /user/bus 	    | User		| POST		| AddNewBusOrder    | MessageResponse   | Trigger table seats.                                          |
| 11 | /user/train	    | User		| GET		| GetTrainList      | TrainsResponse    | -                                                             |
| 12 | /user/train	    | User		| POST		| AddNewTrainOrder  | MessageResponse   | Trigger table seats.                                          |
| 13 | /user/tour	    | User		| GET		| GetTourList       | ToursResponse     | -                                                             |
| 14 | /user/tour	    | User		| POST		| AddNewTourOrder   | MessageResponse   | -                                                             |
| 15 | /user/order	    | User		| GET		| GetUserOrder      | OrdersResponse    | -                                                             |
| 16 | /user/order 	    | User		| DELETE	| RequestRefund     | MessageResponse   | Insert new data to table refunds.                             |
| 17 | /mitra	        | Mitra		| POST		| AddNewMitra	    | MessageResponse   | Register mitra.                                               |
| 18 | /mitra	        | Mitra		| PUT		| UpdateMitra   	| MessageResponse   | -                                                             |
| 19 | /mitra/product   | Mitra		| GET		| GetProductList   	| ProductsResponse  | Query berbeda berdasarkan jenis mitra.                        |
| 20 | /mitra/product   | Mitra		| POST  	| AddNewProduct   	| MessageResponse   | Query berbeda berdasarkan jenis mitra.                        |
| 21 | /mitra/product   | Mitra		| DELETE	| DeleteProduct   	| MessageResponse   | Query berbeda berdasarkan jenis mitra.                        |
| 22 | /admin/refund    | Mitra		| GET		| GetRefundList   	| RefundResponse    | -                                                             |
| 23 | /admin/refund    | Mitra		| DELETE	| ApproveRefund   	| MessageResponse   | -                                                             |


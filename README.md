# Travelokay-Project

| No | Endpoint	        | Access	| Method	| Function                  | Response                          | Desc                                                                                                  |
|----| ---------------- | --------- | --------- | ------------------------- | --------------------------------- | ----------------------------------------------------------------------------------------------------- |
| 1  | /login	        | All		| POST		| Login                     | UserResponse, PartnerResponse     | Saat login berhasil, akan menghasilkan cookie berisi token. Response berbeda tergantung jenis user.   |
| 2  | /logout	        | All		| GET		| Logout                    | MessageResponse                   | Menghapus cookie yang berisi token.                                                                   |
| 3  | /user		    | All		| POST		| AddNewUser	            | MessageResponse                   | Register user.                                                                                        |
| 4  | /user		    | User		| PUT		| UpdateUser	            | MessageResponse                   | -                                                                                                     |
| 5  | /user/hotel	    | User		| GET		| GetHotelList              | HotelsResponse                    | -                                                                                                     |
| 6  | /user/hotel/room | User		| GET		| GetRoomList               | RoomsResponse                     | -                                                                                                     |
| 7  | /user/hotel/room | User		| POST		| AddNewHotelOrder          | HotelOrderResponse                | Trigger update tabel rooms.                                                                           |
| 8  | /user/flight	    | User		| GET		| GetFlightList             | FlightsResponse                   | -                                                                                                     |
| 9  | /user/flight	    | User		| POST		| AddNewFlightOrder         | FlightOrderResponse               | Trigger update tabel seats.                                                                           |
| 10 | /user/bus 	    | User		| GET		| GetBusList                | BusTripsResponse                  | -                                                                                                     |
| 11 | /user/bus 	    | User		| POST		| AddNewBusOrder            | BusOrderResponse                  | Trigger update tabel seats.                                                                           |
| 12 | /user/train	    | User		| GET		| GetTrainList              | TrainTripsResponse                | -                                                                                                     |
| 13 | /user/train	    | User		| POST		| AddNewTrainOrder          | TrainOrderResponse                | Trigger update tabel seats.                                                                           |
| 14 | /user/tour	    | User		| GET		| GetTourList               | ToursResponse                     | -                                                                                                     |
| 15 | /user/tour	    | User		| POST		| AddNewTourOrder           | TourOrderResponse                 | -                                                                                                     |
| 16 | /user/order	    | User		| GET		| GetUserOrder              | UserOrdersResponse                |                                                                                                       |
| 17 | /user/order 	    | User		| DELETE	| RequestRefund             | MessageResponse                   | Update order_status.                                                                                  |
| 18 | /partner	        | All		| POST		| AddNewMitra	            | MessageResponse                   | Register partner.                                                                                     |
| 19 | /partner	        | Partner	| PUT		| UpdateMitra   	        | MessageResponse                   | -                                                                                                     |
| 20 | /partner/flight  | Partner	| GET		| GetFlightPartnerList   	| FlightsResponse                   | Cek apakah perusahaan partner sama dengan params.                                                     |
| 21 | /partner/flight  | Partner	| POST  	| AddNewFlight   	        | FlightResponse                    | Query berbeda berdasarkan jenis partner.                                                              |
| 22 | /partner/flight  | Partner	| DELETE	| DeleteFlight   	        | MessageResponse                   | Query berbeda berdasarkan jenis partner.                                                              |
| 23 | /admin/refund    | Admin		| GET		| GetRefundList   	        | RefundsResponse                   | -                                                                                                     |
| 24 | /admin/refund    | Admin		| DELETE	| ApproveRefund   	        | MessageResponse                   | -                                                                                                     |


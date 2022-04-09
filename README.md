# Travelokay-Project

| No | Endpoint	        | Access	| Method	| Function                  | Response                          | Desc                                                                                                  |
|----| ---------------- | --------- | --------- | ------------------------- | --------------------------------- | ----------------------------------------------------------------------------------------------------- |
| 1  | /login	        | All		| POST		| Login                     | UserResponse, PartnerResponse     | Saat login berhasil, akan menghasilkan cookie berisi token. Response berbeda tergantung jenis user.   |
| 2  | /logout	        | All		| GET		| Logout                    | MessageResponse                   | Menghapus cookie yang berisi token.                                                                   |
| 3  | /user		    | All		| POST		| AddNewUser	            | MessageResponse                   | Register user.                                                                                        |
| 4  | /user		    | User		| PUT		| UpdateUser	            | MessageResponse                   | -                                                                                                     |
| 5  | /user/hotel	    | All		| GET		| GetHotelList              | HotelsResponse                    | -                                                                                                     |
| 6  | /user/hotel/room | All		| GET		| GetRoomList               | RoomsResponse                     | -                                                                                                     |
| 7  | /user/hotel/room | User		| POST		| AddNewRoomOrder          | HotelOrderResponse                | Trigger update tabel rooms.                                                                           |
| 8  | /user/flight	    | All		| GET		| GetFlightList             | FlightsResponse                   | -                                                                                                     |
| 9  | /user/flight	    | User		| POST		| AddNewFlightOrder         | FlightOrderResponse               | Trigger update tabel seats.                                                                           |
| 10 | /user/bus 	    | All		| GET		| GetBusList                | BusTripsResponse                  | -                                                                                                     |
| 11 | /user/bus 	    | User		| POST		| AddNewBusOrder            | BusOrderResponse                  | Trigger update tabel seats.                                                                           |
| 12 | /user/train	    | All		| GET		| GetTrainList              | TrainTripsResponse                | -                                                                                                     |
| 13 | /user/train	    | User		| POST		| AddNewTrainOrder          | TrainOrderResponse                | Trigger update tabel seats.                                                                           |
| 14 | /user/tour	    | All		| GET		| GetTourList               | ToursResponse                     | -                                                                                                     |
| 15 | /user/tour	    | User		| POST		| AddNewTourOrder           | TourOrderResponse                 | -                                                                                                     |
| 16 | /user/order	    | User		| GET		| GetUserOrder              | UserOrdersResponse                |                                                                                                       |
| 17 | /user/order 	    | User		| PUT	    | RequestRefund             | MessageResponse                   | Update order_status.                                                                                  |
| 18 | /partner	        | All		| POST		| AddNewPartner	            | MessageResponse                   | Register partner.                                                                                     |
| 19 | /partner	        | Partner	| PUT		| UpdatePartner   	        | MessageResponse                   | -                                                                                                     |
| 20 | /partner/flight  | Partner	| GET		| GetFlightPartnerList   	| FlightsResponse                   | Cek berdasarkan nama perusahaan partner.                                                              |
| 21 | /partner/flight  | Partner	| POST  	| AddNewFlight   	        | FlightResponse                    | Cek berdasarkan nama perusahaan partner.                                                              |
| 22 | /partner/flight  | Partner	| DELETE	| DeleteFlight   	        | MessageResponse                   | Cek berdasarkan nama perusahaan partner.                                                              |
| 23 | /admin/refund    | Admin		| GET		| GetRefundList   	        | RefundsResponse                   | -                                                                                                     |
| 24 | /admin/refund    | Admin		| DELETE	| ApproveRefund   	        | MessageResponse                   | -                                                                                                     |


-- name: GetRestaurantHistory :many
SELECT id, Date, Time from Visit
    where UserId = ? and RestaurantId = ?;

-- name: GetOrdersForVisit :many
SELECT d.Name, o.Rating, o.ReviewText from
    Orders o join Dish d on o.DishId = d.id
    where o.VisitId = ?;
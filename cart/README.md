# Cart Service

Handles tasks below:

+ Add product to cart,
+ Remove product from cart,
+ Remove all items in cart,
+ List products in cart.

It listens and publishes events via Redis.
After product added to a cart, publishes message to Redis pub/sub for populating product data of added product.

Stores product data and cart data on Mongodb.

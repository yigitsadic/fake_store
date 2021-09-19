# Favourites Service

Favorites

- Service will listen for product information for favorites service event from redis.
- After initializing record, service will fire event to redis to populate record with product information
- Every record will start record with **pending** state.

It should have four endpoints:
1. List favorites - will respond with completed state records -
2. Add to favorites with user ID and product ID
3. Remove from favorites user ID and product ID
4. Product added to favorites service will called with user ID and product ID and return boolean.


# Fake Payment Provider

Mocks PayPal, Stripe like payment provider systems hook and redirect mechanisms.

When request made for create payment intent endpoint, it creates record in database and creates a URL for customer to visit.

In the page that customer visits, there are mock fields for credit card data (ccv, credit card number etc.).

When customer presses "Complete Payment" button, it will send requests to requested hook URL and redirect to
desired page -success or failure.

For simplicity, I didn't implement any authentication nor authorization for workflow.

You were assigned to implement a POS(Point of Sale) system for our client who owns a shop name “Little Brown Book Shop”. They need a precise system, easy to use, and also an elegant user interface.

Their requirements are simple. Firstly, a cashier stands at the POS and use LINE Login to access the POS. After he/she verified their account successfully, the system will show the first page he/she sees is the item-list page. If customers bring books from the shelf and come to checkout, the cashier just has to add what item in the customer's basket, And the system will calculate the total price and some discounts if it meets the criteria. Luckily, there is no tax/vat here. When all customer's item was added the system is ready to proceed to the next step, payment. At the Little Brown, they only accept cash (paying by credit card is still on their mind). In this way, the cashier has to fill the amount of cash into the system and let it process the change back to the customers. Don't forget to show the cashier that everything alright through the payment process. At here, you have to show the customers screen that their payment was succeeded (instead of using paper receipt, save the world!). After everything is done, the system must ready for the next order again.

As we mention about the discount, there are special discount only for all Harry Potter book series (7 books);

buy 2 unique series books discount 10% of those 2 books
buy 3 unique series books discount 11% of those 3 books
buy 4 unique series books discount 12% of those 4 books
buy 5 unique series books discount 13% of those 5 books
buy 6 unique series books discount 14% of those 6 books
buy 7 unique series books discount 15% of those 7 books
For example, assume we have an order;

Harry Potter and the Philosopher's Stone (I). x2 700
Harry Potter and the Chamber of Secrets (II) x1 350
The Fork, the Witch, and the Worm x1 260
Discount 70
Net 1240
You'll see the unique combination of Harry (I), 350฿ and Harry (II), 350฿ will get the discount off 10% (not 2 books of Harry I) which is 70฿.

You, as our best front-end developer, have to develop all of these user-facing parts within 1 week giving our client visibility of the flow they will get after everything is bound together. So, most of the things in this phase will just use only client-side technologies except a client book's API which has been provided by our API team (see the document below).

So, we hope you give them everything essential for this task to impress our client. And don't forget the correctness is a MUST.

However, we have standard practices as follows;

Vue.js Framework only
Write a unit test, please
Reflect on any environment-specific
Apply good UX in mind
Use source control and commit log in a proper way
Document well

API
list books

Beta GET https://api.lbbs.line-beta.me/b/5c52a1be15735a25423d3540
Production GET https://json-bin.netlify.com/books.json
[For Fullstack position]
Instead of using the API above, it required to develop API by using Go language and connect your frontend to this API. Please follow the instructions below:

Record who get accessed to the POS system after a cashier verified themselves successfully with LINE Login
Develop API to list books by using the sample data from https://json-bin.netlify.com/books.json
Develop API to get the discount based on the business in the previous section (frontend assignment) and call this API when you want to get the discount to update your UI.
Write the unit test for your APIs.
Provide API authentication that frontend will use for accessing API (use access token from LINE Login is also okay)
Run your API in Docker Container and provide instructions to start the container in README.
After you have done your work please let us know by;

Write README for how to set up your project and work on local development and send us your source control link
Deploy the project to a server or could and also send us a link, of cause, we won't check your work on our local machine. However, we might deploy your project on our infrastructure (beta environment), be sure that it will work as well
zip that information (maybe just a text file) and upload it here in the remoteinterview.io this action will be a submitting your work. By the way, you can re-upload it anytime
Make sure we can access everything, check it twice!

A bit about mood and tone, if you want to color your interface please consider using our company – LINE color as a guideline. Take a look here: https://design-org.line-apps.com

Note that any additional feature from these requirements is a plus and any impression that you made to us is a very plus plus.

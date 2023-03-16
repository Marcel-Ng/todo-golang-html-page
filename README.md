# Intro
I kind of like started building this thinking to build a todo list put turned out to build a Faq page handler

The application has 
The faqs page where you can 
    * see all the faqs
    * search the faqs
    * pagination
    * Javascript to expand a faq
    * Manage faq  button taking you to the admin page
Admin Page
    * Edit Faq button for each faq
    * Delete Faq button for each faq
    * Add Faq button that takes you to add faq page
Add Faq Page
    * Form that allows you to add a new faq
Edit Faq Page
    * Pre populated form that U can edit the content and update the stuff
/api/1/faqs/0
an api that is exposed via that endpoint and returns a JSON

# Maybe will add the Auth like login or and API auth



# How to run this application

To run this application just run
`go run todo.go`

# To build the Go binary 

Run the command below on the root of the application
`go build`

# TODOs:
- Add the partial files, this will serve as the navigations
- get the styling to work and the javascript and the css to work
- add the following /, /add, /done/{id},

package actions

import (
	"fmt"
	"github.com/constellation-project/goshortie/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Redirection)
// DB Table: Plural (redirections)
// Resource: Plural (Redirections)
// Path: Plural (/redirections)
// View Template Folder: Plural (/templates/redirections/)

// RedirectionsResource is the resource for the Redirection model
type RedirectionsResource struct {
	buffalo.Resource
}

// List gets all Redirections. This function is mapped to the path
// GET /redirections
func (v RedirectionsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	redirections := &models.Redirections{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Redirections from the DB
	if err := q.All(redirections); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, redirections))
}

// Show gets the data for one Redirection. This function is mapped to
// the path GET /redirections/{redirection_id}
func (v RedirectionsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Redirection
	redirection := &models.Redirection{}

	// To find the Redirection the parameter redirection_id is used.
	if err := tx.Find(redirection, c.Param("redirection_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, redirection))
}

// New renders the form for creating a new Redirection.
// This function is mapped to the path GET /redirections/new
func (v RedirectionsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Redirection{}))
}

// Create adds a Redirection to the DB. This function is mapped to the
// path POST /redirections
func (v RedirectionsResource) Create(c buffalo.Context) error {
	// Allocate an empty Redirection
	redirection := &models.Redirection{}

	// Bind redirection to the html form elements
	if err := c.Bind(redirection); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(redirection)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, redirection))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "redirection.created.success"))
	// and redirect to the redirections index page
	return c.Render(201, r.Auto(c, redirection))
}

// Edit renders a edit form for a Redirection. This function is
// mapped to the path GET /redirections/{redirection_id}/edit
func (v RedirectionsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Redirection
	redirection := &models.Redirection{}

	if err := tx.Find(redirection, c.Param("redirection_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, redirection))
}

// Update changes a Redirection in the DB. This function is mapped to
// the path PUT /redirections/{redirection_id}
func (v RedirectionsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Redirection
	redirection := &models.Redirection{}

	if err := tx.Find(redirection, c.Param("redirection_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Redirection to the html form elements
	if err := c.Bind(redirection); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(redirection)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, redirection))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "redirection.updated.success"))
	// and redirect to the redirections index page
	return c.Render(200, r.Auto(c, redirection))
}

// Destroy deletes a Redirection from the DB. This function is mapped
// to the path DELETE /redirections/{redirection_id}
func (v RedirectionsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Redirection
	redirection := &models.Redirection{}

	// To find the Redirection the parameter redirection_id is used.
	if err := tx.Find(redirection, c.Param("redirection_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(redirection); err != nil {
		return err
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "redirection.destroyed.success"))
	// Redirect to the redirections index page
	return c.Render(200, r.Auto(c, redirection))
}

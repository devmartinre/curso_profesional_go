package CodigoFacilito

func (c *Curso) getTitulo() string {
	return c.Titulo
}

func (c *Curso) ObtenerTitulo() string {
	return c.getTitulo()
}

import { Navbar, Nav, Container } from 'react-bootstrap'
import { NavLink } from 'react-router-dom'

const links = [
  { to: '/', label: 'Dashboard' },
  { to: '/restaurantes', label: 'Restaurantes' },
  { to: '/ordenes', label: 'Órdenes' },
  { to: '/usuarios', label: 'Usuarios' },
  { to: '/resenas', label: 'Reseñas' },
]

export default function NavigationBar() {
  return (
    <Navbar
      expand="lg"
      style={{ backgroundColor: '#ea580c' }} // naranja tipo orange-600
      variant="dark"
      className="shadow-sm"
    >
      <Container>
        <Navbar.Brand className="fw-bold text-white">
          🍽 RestaurantApp
        </Navbar.Brand>

        <Navbar.Toggle aria-controls="main-navbar" />

        <Navbar.Collapse id="main-navbar">
          <Nav className="ms-4">
            {links.map((l) => (
              <Nav.Link
                key={l.to}
                as={NavLink}
                to={l.to}
                end={l.to === '/'}
                style={{ textDecoration: 'none' }}
                className="text-white px-3"
              >
                {l.label}
              </Nav.Link>
            ))}
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

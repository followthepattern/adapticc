import { Link } from 'react-router-dom'
 
export default function NotFound() {
  return (
    <div>
      <h2>Not Found</h2>
      <p>Could not find requested resource</p>
      <p>
        View <Link to="/users">all posts</Link>
      </p>
    </div>
  )
}
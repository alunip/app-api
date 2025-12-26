import './App.css'
import { useConfig } from './hooks/useConfig'

function App() {
  const { data, loading, error, refetch } = useConfig()

  if (loading) {
    return (
      <div className="app">
        <h1>Loading configuration...</h1>
      </div>
    )
  }

  if (error) {
    return (
      <div className="app">
        <h1>Error loading configuration</h1>
        <p style={{ color: 'red' }}>{error.message}</p>
        <button onClick={refetch}>Retry</button>
      </div>
    )
  }

  if (!data) {
    return (
      <div className="app">
        <h1>No configuration data available</h1>
        <button onClick={refetch}>Retry</button>
      </div>
    )
  }

  return (
    <div className="app">
      <h1>Application Configuration</h1>
      <div className="config-info">
        <p><strong>Name:</strong> {data.name}</p>
        <p><strong>Version:</strong> {data.version}</p>
      </div>
      <button onClick={refetch}>Refresh Configuration</button>
    </div>
  )
}

export default App

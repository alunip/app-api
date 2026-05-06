import './App.css'
import { useTranslation } from 'react-i18next'
import { useConfig } from './hooks/useConfig'

function App() {
  const { t } = useTranslation()
  const { data, loading, error, refetch } = useConfig()

  if (loading) {
    return (
      <div className="app">
        <h1>{t('loading')}</h1>
      </div>
    )
  }

  if (error) {
    return (
      <div className="app">
        <h1>{t('error.title')}</h1>
        <p style={{ color: 'red' }}>{error.message}</p>
        <button onClick={refetch}>{t('error.retry')}</button>
      </div>
    )
  }

  if (!data) {
    return (
      <div className="app">
        <h1>{t('noData.title')}</h1>
        <button onClick={refetch}>{t('error.retry')}</button>
      </div>
    )
  }

  return (
    <div className="app">
      <h1>{t('config.title')}</h1>
      <div className="config-info">
        <p><strong>{t('config.name')}:</strong> {data.name}</p>
        <p><strong>{t('config.version')}:</strong> {data.version}</p>
      </div>
      <button onClick={refetch}>{t('config.refresh')}</button>
    </div>
  )
}

export default App

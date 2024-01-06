import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LandingPage from './pages/Landing';
import NotFoundPage from './pages/NotFound';
import LinkShortenPage from './pages/LinkShorten';
import RedirectionPage from './pages/Redirection';
import LinkGeneratedPage from './pages/LinkGenerated';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<LandingPage />} />
        <Route
          path="/statisticlink/:link_id/:target"
          element={<RedirectionPage />}
        />
        <Route path="/shorten" element={<LinkShortenPage />} />
        <Route path="/generated/:link" element={<LinkGeneratedPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

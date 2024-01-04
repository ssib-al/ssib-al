import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LandingPage from './pages/Landing';
import NotFoundPage from './pages/NotFound';
import LinkShortenPage from './pages/LinkShorten';
import RedirectionPage from './pages/Redirection';

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
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

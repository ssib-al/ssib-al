import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LandingPage from './pages/Landing';
import NotFound from './pages/NotFound';
import LinkShortenPage from './pages/LinkShorten';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<LandingPage />} />
        <Route path="/shorten" element={<LinkShortenPage />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

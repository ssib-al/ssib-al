import { useParams } from 'react-router-dom';
import NotFoundPage from './NotFound';

const RedirectionPage = () => {
  const { link_id, target } = useParams();
  if (!link_id || !target) {
    return <NotFoundPage />;
  }
  return (
    <div>
      <p>link-id: {link_id}</p>
      <p>target: {target}</p>
    </div>
  );
};

export default RedirectionPage;

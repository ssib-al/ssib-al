import { useParams } from 'react-router-dom';
import Navbar from '../components/Navbar';
import NotFoundPage from './NotFound';
import {
  ArrowPathIcon,
  CheckCircleIcon,
  CloudArrowUpIcon,
  FingerPrintIcon,
  LinkIcon,
  LockClosedIcon,
} from '@heroicons/react/24/outline';
import Notification from '../components/Notification';
import { useState } from 'react';
import ShareModal from '../components/ShareModal';

const features = [
  {
    name: 'Push to deploy',
    description:
      'Morbi viverra dui mi arcu sed. Tellus semper adipiscing suspendisse semper morbi. Odio urna massa nunc massa.',
    icon: CloudArrowUpIcon,
  },
  {
    name: 'SSL certificates',
    description:
      'Sit quis amet rutrum tellus ullamcorper ultricies libero dolor eget. Sem sodales gravida quam turpis enim lacus amet.',
    icon: LockClosedIcon,
  },
  {
    name: 'Simple queues',
    description:
      'Quisque est vel vulputate cursus. Risus proin diam nunc commodo. Lobortis auctor congue commodo diam neque.',
    icon: ArrowPathIcon,
  },
  {
    name: 'Advanced security',
    description:
      'Arcu egestas dolor vel iaculis in ipsum mauris. Tincidunt mattis aliquet hac quis. Id hac maecenas ac donec pharetra eget.',
    icon: FingerPrintIcon,
  },
];

const LinkGeneratedPage = () => {
  const [linkShareModalOpen, setLinkShareModalOpen] = useState(false);
  const [linkCopiedNotificationOpen, setLinkCopiedNotificationOpen] =
    useState(false);
  const { link } = useParams();
  if (!link) {
    return <NotFoundPage />;
  }
  return (
    <div className="bg-white">
      <Navbar />
      <section className="relative isolate h-full overflow-hidden bg-white px-6 py-24 sm:py-32 lg:px-8">
        <div className="absolute inset-0 -z-10 bg-[radial-gradient(45rem_50rem_at_top,theme(colors.indigo.100),white)] opacity-20" />
        <div className="absolute inset-y-0 right-1/2 -z-10 mr-16 w-[200%] origin-bottom-left skew-x-[-30deg] bg-white shadow-xl shadow-indigo-600/10 ring-1 ring-indigo-50 sm:mr-28 lg:mr-0 xl:mr-16 xl:origin-center" />
        <div className="mx-auto max-w-2xl pt-48 lg:max-w-4xl">
          <h2 className=" text-center text-3xl font-bold text-indigo-600 sm:text-4xl">
            링크가 생성되었습니다
          </h2>
          <figure className="mt-10">
            <blockquote className="text-center">
              <LinkIcon className="-mt-1 mr-2 inline-block h-6 w-6 text-gray-800" />
              <button
                className="font-regular text-xl leading-8 text-gray-700 hover:underline sm:text-2xl sm:leading-9"
                onClick={() => {
                  navigator.clipboard.writeText(link);
                  setLinkCopiedNotificationOpen(true);
                }}
              >
                {link}
              </button>
            </blockquote>
            <p className="mt-3 text-center text-base text-sm leading-6 text-gray-600">
              위의 링크를 눌러 복사하실 수 있습니다
            </p>
            <div className="mx-auto mt-12 w-auto text-center">
              <button
                type="button"
                className="rounded-md bg-indigo-600 px-5 py-3 text-base font-medium text-white hover:bg-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:px-8"
                onClick={async () => {
                  if (/Mobi|Android/i.test(navigator.userAgent)) {
                    await navigator.share({
                      title: '씨발 - 링크단축 공유',
                      text: '공유된 링크를 확인하세요!',
                      url: link,
                    });
                  } else {
                    setLinkShareModalOpen(true);
                  }
                }}
              >
                링크 공유하기
              </button>
              <a
                className="ml-6 text-base font-medium text-gray-700 sm:ml-9 sm:pl-2 sm:pr-3"
                href="/"
              >
                메인으로 돌아가기 <span aria-hidden="true">&rarr;</span>
              </a>
            </div>
          </figure>
        </div>
        <div className="mt-32 py-24 sm:py-32">
          <div className="mx-auto max-w-7xl px-6 lg:px-8">
            <div className="mx-auto max-w-2xl lg:text-center">
              <h2 className="font-regular text-base leading-7 text-indigo-600">
                다양한 기능을 사용해보고 싶으신가요?
              </h2>
              <p className="mt-2 text-2xl font-semibold tracking-tight text-gray-900 sm:text-3xl">
                패스워드 또는 계정을 사용해 링크를 관리해보세요
              </p>
              <p className="mt-6 text-lg leading-8 text-gray-600">
                Quis tellus eget adipiscing convallis sit sit eget aliquet quis.
                Suspendisse eget egestas a elementum pulvinar et feugiat blandit
                at. In mi viverra elit nunc.
              </p>
            </div>
            <div className="mx-auto mt-16 max-w-2xl sm:mt-20 lg:mt-24 lg:max-w-4xl">
              <dl className="grid max-w-xl grid-cols-1 gap-x-8 gap-y-10 lg:max-w-none lg:grid-cols-2 lg:gap-y-16">
                {features.map((feature) => (
                  <div key={feature.name} className="relative pl-16">
                    <dt className="text-base font-semibold leading-7 text-gray-900">
                      <div className="absolute left-0 top-0 flex h-10 w-10 items-center justify-center rounded-lg bg-indigo-600">
                        <feature.icon
                          className="h-6 w-6 text-white"
                          aria-hidden="true"
                        />
                      </div>
                      {feature.name}
                    </dt>
                    <dd className="mt-2 text-base leading-7 text-gray-600">
                      {feature.description}
                    </dd>
                  </div>
                ))}
              </dl>
            </div>
          </div>
        </div>
      </section>
      <ShareModal
        title="링크 공유하기"
        description={link}
        open={linkShareModalOpen}
        setOpen={setLinkShareModalOpen}
        link={link}
      />
      ;
      <Notification
        open={linkCopiedNotificationOpen}
        setOpen={setLinkCopiedNotificationOpen}
        icon={
          <CheckCircleIcon
            className="h-6 w-6 text-green-600"
            aria-hidden="true"
          />
        }
        title="복사 완료!"
        description="복사된 링크를 공유하세요!"
        ttl={3000}
      />
    </div>
  );
};

export default LinkGeneratedPage;

import { Dialog, Transition } from '@headlessui/react';
import SsibalLogo from '../assets/ssibal-logo.svg?react';
import Navbar from '../components/Navbar';
import Notification from '../components/Notification';
import { Fragment, useRef, useState } from 'react';
import {
  ExclamationCircleIcon,
  QuestionMarkCircleIcon,
} from '@heroicons/react/24/outline';

const urlRegex: RegExp = new RegExp(
  '^[-a-zA-Z0-9@:%._+~#=]{2,256}.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_+.~#?&//=]*)$',
);

const LandingPage = () => {
  const [linkModalOpen, setLinkModalOpen] = useState(false);
  const [linkInvalidNotificationOpen, setLinkInvalidNotificationOpen] =
    useState(false);

  const cancelButtonRef = useRef(null);
  const linkStrInputRef = useRef<HTMLInputElement>(null);
  const linkSSLInputRef = useRef<HTMLSelectElement>(null);
  const openLinkModal = () => {
    if (!linkStrInputRef.current) return;
    if (!urlRegex.test(linkStrInputRef.current.value)) {
      setLinkInvalidNotificationOpen(true);
      return;
    }
    setLinkInvalidNotificationOpen(false);
    setLinkModalOpen(true);
  };
  const continueToLinkGenPage = () => {
    setLinkModalOpen(false);
    if (!linkStrInputRef.current || !linkSSLInputRef.current) return;
    const params = new URLSearchParams({
      ssl: linkSSLInputRef.current.value == 'https://' ? 'true' : 'false',
      link: linkStrInputRef.current.value,
    });
    window.location.href = '/shorten?' + params.toString();
  };
  return (
    <div className="bg-white">
      <Navbar />
      <div className="relative isolate overflow-hidden bg-gradient-to-b from-indigo-100/20 pt-14">
        <div
          className="absolute inset-y-0 right-1/2 -z-10 -mr-96 w-[200%] origin-top-right skew-x-[-30deg] bg-white shadow-xl shadow-indigo-600/10 ring-1 ring-indigo-50 sm:-mr-80 lg:-mr-96"
          aria-hidden="true"
        />
        <div className="mx-auto max-w-7xl px-6 py-32 sm:py-40 lg:px-8">
          <div className="mx-auto max-w-2xl lg:mx-0 lg:grid lg:max-w-none lg:grid-cols-2 lg:gap-x-16 lg:gap-y-6 xl:grid-cols-1 xl:grid-rows-1 xl:gap-x-8">
            <h1 className="racking-tight max-w-2xl text-4xl font-bold text-gray-900 sm:text-6xl sm:leading-tight lg:col-span-2 xl:col-auto">
              <span className="text-3xl font-normal">
                한 번 들으면 <span className="font-semibold">잊을 수 없는</span>
              </span>
              <br></br>
              링크단축 서비스 "씨발"
            </h1>
            <div className="mt-6 max-w-xl lg:mt-0 xl:col-end-1 xl:row-start-1">
              <p className="text-base text-gray-500">
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Quas
                cupiditate laboriosam fugiat.
              </p>
              <div className="pt-5">
                <label
                  htmlFor="phone-number"
                  className="block text-sm font-medium leading-6 text-gray-900"
                >
                  단축할 링크를 입력하세요
                </label>
                <div className="relative mt-2 rounded-md shadow-sm">
                  <div className="absolute inset-y-0 left-0 flex items-center">
                    <select
                      ref={linkSSLInputRef}
                      id="ssl"
                      name="ssl"
                      className="h-full rounded-md border-0 bg-transparent py-0 pl-3 pr-7 text-gray-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm"
                    >
                      <option>http://</option>
                      <option>https://</option>
                    </select>
                  </div>
                  <input
                    ref={linkStrInputRef}
                    type="text"
                    name="link"
                    id="link"
                    className="block w-full rounded-md border-0 py-1.5 pl-20 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                    placeholder="www.example.com"
                  />
                </div>
              </div>
              <div className="mt-10 flex items-center gap-x-6">
                <button
                  onClick={() => openLinkModal()}
                  className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  단축하기
                </button>
                <a
                  href="#"
                  className="text-sm font-semibold leading-6 text-gray-900"
                >
                  더 알아보기 <span aria-hidden="true">→</span>
                </a>
              </div>
            </div>
            <SsibalLogo
              width={400}
              height={400}
              className="lg:display mt-10 hidden aspect-[6/5] w-full max-w-lg rounded-2xl object-cover drop-shadow-2xl sm:mt-16 lg:mt-0 lg:block lg:max-w-none xl:row-span-2 xl:row-end-2 xl:mt-36"
            />
          </div>
        </div>
        <div className="absolute inset-x-0 bottom-0 -z-10 h-24 bg-gradient-to-t from-white sm:h-32" />
      </div>
      <Transition.Root show={linkModalOpen} as={Fragment}>
        <Dialog
          as="div"
          className="relative z-10"
          initialFocus={cancelButtonRef}
          onClose={setLinkModalOpen}
        >
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
          </Transition.Child>

          <div className="fixed inset-0 z-10 w-screen overflow-y-auto">
            <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
              <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                enterTo="opacity-100 translate-y-0 sm:scale-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100 translate-y-0 sm:scale-100"
                leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              >
                <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg sm:p-6">
                  <div>
                    <div className="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-green-100">
                      <QuestionMarkCircleIcon
                        className="h-6 w-6 text-green-600"
                        aria-hidden="true"
                      />
                    </div>
                    <div className="mt-3 text-center sm:mt-5">
                      <Dialog.Title
                        as="h3"
                        className="text-base font-semibold leading-6 text-gray-900"
                      >
                        링크 관리가 필요하신가요?
                      </Dialog.Title>
                      <div className="mt-2">
                        <p className="text-sm text-gray-500">
                          링크 관리 수단을 추가하면 통계를 확인하고 링크를 수정
                          및 삭제할 수 있습니다.
                        </p>
                      </div>
                    </div>
                  </div>
                  <div className="mt-5 sm:mt-6 sm:grid sm:grid-flow-row-dense sm:grid-cols-2 sm:gap-3">
                    <button
                      type="button"
                      className="inline-flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 sm:col-start-2"
                      onClick={() => continueToLinkGenPage()}
                    >
                      링크 관리 수단 추가
                    </button>
                    <button
                      type="button"
                      className="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:col-start-1 sm:mt-0"
                      onClick={() => setLinkModalOpen(false)}
                      ref={cancelButtonRef}
                    >
                      스킵하기
                    </button>
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition.Root>
      <Notification
        open={linkInvalidNotificationOpen}
        setOpen={setLinkInvalidNotificationOpen}
        icon={
          <ExclamationCircleIcon
            className="h-6 w-6 text-red-600"
            aria-hidden="true"
          />
        }
        title="링크가 유효하지 않습니다."
        description="링크를 다시 확인해주세요."
      />
    </div>
  );
};

export default LandingPage;

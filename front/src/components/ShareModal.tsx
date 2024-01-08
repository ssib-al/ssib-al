import { Fragment, useEffect } from 'react';
import { Dialog, Transition } from '@headlessui/react';
import { ShareIcon } from '@heroicons/react/24/outline';
import XTwitterIcon from '../assets/x-twitter.svg?react';
import KakaotalkIcon from '../assets/kakaotalk.svg?react';
import { ShareKakao, ShareTwitter } from '../utils/snsShare';

interface ShareModalProps {
  title: string;
  description: string;
  link: string;
  open: boolean;
  setOpen: (open: boolean) => void;
}

const ShareModal: React.FC<ShareModalProps> = (props) => {
  useEffect(() => {
    if (
      document.querySelector(
        `script[src="https://t1.kakaocdn.net/kakao_js_sdk/2.6.0/kakao.min.js"]`,
      )
    )
      return;

    const script = document.createElement('script');
    script.src = 'https://t1.kakaocdn.net/kakao_js_sdk/2.6.0/kakao.min.js';
    script.async = false;
    script.integrity =
      'sha384-6MFdIr0zOira1CHQkedUqJVql0YtcZA1P0nbPrQYJXVJZUkTk/oX4U9GhUIs3/z8';
    script.crossOrigin = 'anonymous';
    document.body.appendChild(script);
  }, []);
  return (
    <>
      <Transition.Root show={props.open} as={Fragment}>
        <Dialog as="div" className="relative z-10" onClose={props.setOpen}>
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
                      <ShareIcon
                        className="h-6 w-6 text-green-600"
                        aria-hidden="true"
                      />
                    </div>
                    <div className="mt-3 text-center sm:mt-5">
                      <Dialog.Title
                        as="h3"
                        className="text-base font-semibold leading-6 text-gray-900"
                      >
                        {props.title}
                      </Dialog.Title>
                      <div className="mt-2">
                        <p className="text-sm text-gray-500">
                          {props.description}
                        </p>
                      </div>
                    </div>
                  </div>
                  {/* share to kakaotalk and twitter */}
                  <div className="mt-5 flex justify-center gap-x-4">
                    <button
                      type="button"
                      className="inline-flex justify-center rounded-md border border-transparent bg-yellow-400 px-4 py-2 text-sm font-medium text-white hover:bg-yellow-500"
                      onClick={() => {
                        ShareKakao(props.link);
                      }}
                    >
                      <KakaotalkIcon className="h-5 w-5 rounded-sm" />
                      <span className="rounded-sm pl-2">카카오톡 공유</span>
                    </button>
                    <button
                      type="button"
                      className="inline-flex justify-center rounded-md border border-transparent bg-gray-800 px-4 py-2 text-sm font-medium text-white hover:bg-gray-900"
                      onClick={() => {
                        ShareTwitter(props.link);
                      }}
                    >
                      <XTwitterIcon className="h-5 w-5 rounded-sm" />
                      <span className="pl-2">트위터 공유</span>
                    </button>
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition.Root>
    </>
  );
};

export default ShareModal;

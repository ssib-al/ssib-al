import { Transition } from '@headlessui/react';
import { XMarkIcon } from '@heroicons/react/16/solid';
import { Fragment, useEffect } from 'react';

interface NotificationProps {
  open: boolean;
  setOpen: (open: boolean) => void;
  icon: React.ReactNode;
  title: string;
  description: string;
  ttl?: number;
}

const Notification: React.FC<NotificationProps> = (props) => {
  useEffect(() => {
    if (props.open) {
      setTimeout(() => {
        props.setOpen(false);
      }, props.ttl || 3000);
    }
  });
  return (
    <div
      aria-live="assertive"
      className="pointer-events-none static fixed inset-0 z-50 flex items-end px-4 py-6 sm:p-6"
    >
      <div className="sm:item-center flex w-full flex-col items-start space-y-4">
        {/* Notification panel, dynamically insert this into the live region when it needs to be displayed */}
        <Transition
          show={props.open}
          as={Fragment}
          enter="transform ease-out duration-300 transition"
          enterFrom="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
          enterTo="translate-y-0 opacity-100 sm:translate-x-0"
          leave="transition ease-in duration-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="pointer-events-auto w-full max-w-sm overflow-hidden rounded-lg bg-white shadow-lg ring-1 ring-black ring-opacity-5">
            <div className="p-4">
              <div className="flex items-start">
                <div className="flex-shrink-0">{props.icon}</div>
                <div className="ml-3 w-0 flex-1 pt-0.5">
                  <p className="text-sm font-medium text-gray-900">
                    {props.title}
                  </p>
                  <p className="mt-1 text-sm text-gray-500">
                    {props.description}
                  </p>
                </div>
                <div className="ml-4 flex flex-shrink-0">
                  <button
                    type="button"
                    className="inline-flex rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                    onClick={() => {
                      props.setOpen(false);
                    }}
                  >
                    <span className="sr-only">Close</span>
                    <XMarkIcon className="h-5 w-5" aria-hidden="true" />
                  </button>
                </div>
              </div>
            </div>
          </div>
        </Transition>
      </div>
    </div>
  );
};

export default Notification;

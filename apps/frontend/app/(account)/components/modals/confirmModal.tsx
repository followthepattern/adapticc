import classNames from 'classnames';
import React, { MouseEventHandler } from 'react';
import { Button, Dialog, DialogTrigger, Modal, ModalOverlay } from 'react-aria-components';
import TertiaryButton from '../buttons/tertiaryButton';
import SecondaryButton from '../buttons/secondaryButton';
import PrimaryButton from '../buttons/primaryButton';

interface ConfirmModalProperties {
  className?: string
  children?: any
  title: string
  body?: string
  onConfirm?: MouseEventHandler<HTMLButtonElement>
}

export default function ConfirmModal(props: ConfirmModalProperties) {
  return (
    <DialogTrigger>
      <Button className={classNames(props.className, TertiaryButton.ClassName)}>
        {props.children}
      </Button>

      <ModalOverlay
        className={({ isEntering, isExiting }) =>
          classNames("fixed inset-0 z-50 overflow-y-auto bg-black/25 flex min-h-full items-end sm:items-center justify-center p-4 text-center backdrop-blur",
            {
              "animate-in fade-in duration-300 ease-out": isEntering,
              "animate-out fade-out duration-200 ease-in": isExiting,
            })}
      >
        <Modal
          className={({ isEntering, isExiting }) =>
            classNames("w-full max-w-md overflow-hidden rounded-2xl bg-white p-6 text-left shadow-xl",
              {
                "animate-in zoom-in-95 ease-out duration-300": isEntering,
                "animate-out zoom-out-95 ease-in duration-200": isExiting
              })}
        >
          <Dialog className="outline-none">
            {({ close }) => (
              <>
                <div
                  slot="title"
                  className="my-0 text-xl font-semibold"
                >
                  {props.title}
                </div>
                <p className="mt-3 text-gray-700">
                  {props.body}
                </p>
                <div className="flex flex-col-reverse gap-2 mt-6 sm:flex-row sm:flex sm:justify-start">
                  <PrimaryButton
                    onClick={props.onConfirm}
                  >
                    Delete
                  </PrimaryButton>
                  <SecondaryButton
                    onClick={close}

                  >
                    Cancel
                  </SecondaryButton>
                </div>
              </>
            )}
          </Dialog>
        </Modal>
      </ModalOverlay>
    </DialogTrigger >
  );
}
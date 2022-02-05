import create, { GetState, SetState } from "zustand";
import { devtools, persist, StoreApiWithPersist } from "zustand/middleware";

const storeNames = {
  navbar: "navbar-store",
};

interface NavbarExpandedSate {
  navbarExpanded: boolean;
  toggleNavbarExpanded: () => void;
};

export const useNavbarStore = create<
  NavbarExpandedSate,
  SetState<NavbarExpandedSate>,
  GetState<NavbarExpandedSate>,
  StoreApiWithPersist<NavbarExpandedSate>
>(
    persist(devtools(
        (set) => ({
            navbarExpanded: true,
            toggleNavbarExpanded: () => set((state) => ({navbarExpanded: !state.navbarExpanded})),
        })
    ),{
        name: storeNames.navbar,
        getStorage: () => localStorage
    })
)
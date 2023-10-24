import { getContext, hasContext, setContext } from "svelte";
import { readable, writable } from "svelte/store";

// export const useStore =  <T, A>(
//     name: string,
//     fn: () => T,
// ):T => {
//     if (hasContext(name)) {
//         return getContext<T>(name);
//     }
//     const _value = fn();
//     setContext(name, _value);
//     return _value;
// };
//
// export const useWritable = <T>(name: string):T =>
//     useStore(name, writable<T>);
//
// export const useReadable = <T>(name: string):T =>
//     useStore(name, readable);
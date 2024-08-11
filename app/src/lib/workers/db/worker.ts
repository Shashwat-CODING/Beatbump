/* eslint-disable @typescript-eslint/no-explicit-any */
/// <reference lib="webworker" />
import { dbHandler } from "./db";
import type { Actions, Methods } from "./types";

const worker = self as unknown as DedicatedWorkerGlobalScope;

worker.onmessage = async <
	Action extends Actions = any,
	Type extends "favorite" | "playlist" | "playlists" | "favorites" = any,
	Key extends keyof Methods & `${Action}${Capitalize<Type>}` = keyof Methods &
		`${Action & string}${Capitalize<Type>}`,
	Fn extends Methods[Key] = Methods[Key],
>(
	event: MessageEvent<{
		action: Action;
		type: Type;
		params: [...Parameters<Fn>];
	}>,
) => {
	const { action, type, params } = event.data;
	await dbHandler<Action, Type, Key, Fn>(action, type, ...params)
		.then((result) => postMessage(result))
		.catch((err) => {
			console.error(`[DB:TX Error] `, err);
			return err;
		});
};

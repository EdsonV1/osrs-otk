type DynamicRoutes = {
	"/skills/[skill_name]": { skill_name: string }
};

type Layouts = {
	"/": { skill_name?: string };
	"/contact": undefined;
	"/skills": { skill_name?: string };
	"/skills/[skill_name]": { skill_name: string };
	"/tools": undefined;
	"/tools/ardy-knights": undefined;
	"/tools/birdhouses": undefined;
	"/tools/gotr": undefined;
	"/tools/wintertodt": undefined
};

export type RouteId = "/" | "/contact" | "/skills" | "/skills/[skill_name]" | "/tools" | "/tools/ardy-knights" | "/tools/birdhouses" | "/tools/gotr" | "/tools/wintertodt";

export type RouteParams<T extends RouteId> = T extends keyof DynamicRoutes ? DynamicRoutes[T] : Record<string, never>;

export type LayoutParams<T extends RouteId> = Layouts[T] | Record<string, never>;

export type Pathname = "/" | "/contact" | "/skills" | `/skills/${string}` & {} | "/tools" | "/tools/ardy-knights" | "/tools/birdhouses" | "/tools/gotr" | "/tools/wintertodt";

export type ResolvedPathname = `${"" | `/${string}`}${Pathname}`;

export type Asset = "/favicon.png";
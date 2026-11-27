import { InMemUuidGen } from "@/modules/shared/infra/in-mem-uuidGen";
import { Uuids } from "@/modules/shared/domain/uuid";
import { AuthGateway } from "@/modules/auth/domain/authGateway";
import { SessionEvents } from "@/modules/auth/domain/sessionEvents";
import { InMemAuthGateway } from "@/modules/auth/infra/in-mem-authGateway";
import { HttpAuthGateway } from "@/modules/auth/infra/http-authGateway";
import { HttpWorkspaces } from "@/modules/auth/infra/http-workspaces";
import { InMemWorkspaces } from "@/modules/auth/infra/in-mem-workspaces";
import { createHttpClient } from "@/modules/auth/infra/http-client";
import { InMemTokenStorage } from "@/modules/auth/infra/in-mem-tokenStorage";
import { LocalStorageTokenStorage } from "@/modules/auth/infra/localStorage-tokenStorage";
import { AuthSrv } from "@/modules/auth/app/authSrv";
import { WorkspaceSrv } from "@/modules/auth/app/workspaceSrv";
import { Ingredients } from "@/modules/recipe/domain/ingredients";
import { InMemIngredients } from "@/modules/recipe/infra/in-mem-ingredients";
import { IngredientSrv } from "@/modules/recipe/app/ingredientSrv";
import { Ingredient } from "@/modules/recipe/domain/Ingredient";
import { IngredientPlain } from "@/modules/recipe/init/ingredient.init";

export type Dependencies = {
  uuidGen: Uuids;
  authGateway: AuthGateway;
  authSrv: AuthSrv;
  workspaceSrv: WorkspaceSrv;
  ingredients: Ingredients;
  ingredientSrv: IngredientSrv;
};

export const createAppDependencies = (): Dependencies => {
  const sessionEvents = new SessionEvents();
  const tokenStorage = new LocalStorageTokenStorage();
  // Logging out in one tab logs out every tab
  tokenStorage.onRemovedInOtherTab(() => sessionEvents.emitSessionExpired());

  // One shared client, so every module shares the same single-flight refresh
  const http = createHttpClient({
    tokenStorage,
    onSessionExpired: () => sessionEvents.emitSessionExpired(),
  });

  const uuidGen = new InMemUuidGen();
  const authGateway = new HttpAuthGateway(http);
  const ingredients = new InMemIngredients(uuidGen, []);

  return {
    uuidGen,
    authGateway,
    authSrv: new AuthSrv(authGateway, tokenStorage, sessionEvents),
    workspaceSrv: new WorkspaceSrv(new HttpWorkspaces(http)),
    ingredients,
    ingredientSrv: new IngredientSrv(ingredients),
  };
};

// ======== TEST =========================================
// =======================================================

export type TestDependenciesOptions = {
  isUsersSlow?: boolean;
  isResponseError?: boolean;
  isUserUnverified?: boolean;
  isWrongPassword?: boolean;
  isSessionExpired?: boolean;
  initPlainIngredients?: IngredientPlain[];
};

export type TestDependencies = {
  sessionEvents: SessionEvents;
  tokenStorage: InMemTokenStorage;
  uuidGen: InMemUuidGen;
  authGateway: InMemAuthGateway;
  authSrv: AuthSrv;
  workspaces: InMemWorkspaces;
  workspaceSrv: WorkspaceSrv;
  ingredients: InMemIngredients;
  ingredientSrv: IngredientSrv;
};

export const createTestDependencies = (
  opts: TestDependenciesOptions = {},
): TestDependencies => {
  const sessionEvents = new SessionEvents();
  const tokenStorage = new InMemTokenStorage();
  const uuidGen = new InMemUuidGen();
  const authGateway = new InMemAuthGateway(uuidGen, !!opts.isUsersSlow)
    .setIsResponseError(!!opts.isResponseError)
    .setIsUserUnverified(!!opts.isUserUnverified)
    .setIsWrongPassword(!!opts.isWrongPassword)
    .setIsSessionExpired(!!opts.isSessionExpired);
  const workspaces = new InMemWorkspaces(uuidGen);
  const ingredients = new InMemIngredients(
    uuidGen,
    (opts.initPlainIngredients ?? []).map(
      (ip) => new Ingredient(ip.id, ip.name, ip.baseUnit, ip.workspaceId),
    ),
  );

  return {
    sessionEvents,
    tokenStorage,
    uuidGen,
    authGateway,
    authSrv: new AuthSrv(authGateway, tokenStorage, sessionEvents),
    workspaces,
    workspaceSrv: new WorkspaceSrv(workspaces),
    ingredients,
    ingredientSrv: new IngredientSrv(ingredients),
  };
};

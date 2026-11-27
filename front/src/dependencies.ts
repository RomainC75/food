import { Uuids } from "@/modules/shared/domain/uuid";
import { InMemUuidGen } from "@/modules/shared/infra/in-mem-uuidGen";
import { Users } from "@/modules/auth/domain/users";
import { InMemUser } from "@/modules/auth/infra/in-mem-users";
import { HttpUsers } from "@/modules/auth/infra/http-users";
import { AuthSrv } from "@/modules/auth/app/authSrv";
import { Ingredients } from "@/modules/recipe/domain/ingredients";
import { InMemIngredients } from "@/modules/recipe/infra/in-mem-ingredients";
import { IngredientSrv } from "@/modules/recipe/app/ingredientSrv";
import { Ingredient } from "@/modules/recipe/domain/Ingredient";
import { IngredientPlain } from "@/modules/recipe/init/ingredient.init";
import { InMemTokenStorage } from "./modules/auth/infra/in-mem-tokenStorage";
import { LocalStorageTokenStorage } from "./modules/auth/infra/localStorage-tokenStorage";

export type Dependencies = {
  uuidGen: Uuids;
  users: Users;
  authSrv: AuthSrv;
  ingredients: Ingredients;
  ingredientSrv: IngredientSrv;
};

export const createAppDependencies = (): Dependencies => {
  const tokenStorage = new LocalStorageTokenStorage();
  const uuidGen = new InMemUuidGen();
  const users = new HttpUsers(tokenStorage);
  const ingredients = new InMemIngredients(uuidGen, []);

  return {
    uuidGen,
    users,
    ingredients,
    authSrv: new AuthSrv(users, tokenStorage),
    ingredientSrv: new IngredientSrv(ingredients),
  };
};

// ======== TEST =========================================
// =======================================================

export type TestDependenciesOptions = {
  isUsersSlow?: boolean;
  isResponseError?: boolean;
  isUserUnverified?: boolean;
  initPlainIngredients?: IngredientPlain[];
};

export type TestDependencies = {
  tokenStorage: InMemTokenStorage;
  uuidGen: InMemUuidGen;
  users: InMemUser;
  authSrv: AuthSrv;
  ingredients: InMemIngredients;
  ingredientSrv: IngredientSrv;
};

export const createTestDependencies = (
  opts: TestDependenciesOptions = {},
): TestDependencies => {
  const tokenStorage = new InMemTokenStorage();
  const uuidGen = new InMemUuidGen();
  const users = new InMemUser(uuidGen, !!opts.isUsersSlow)
    .setIsResponseError(!!opts.isResponseError)
    .setIsUserUnverified(!!opts.isUserUnverified);
  const ingredients = new InMemIngredients(
    uuidGen,
    (opts.initPlainIngredients ?? []).map(
      (ip) => new Ingredient(ip.id, ip.name, ip.baseUnit, ip.workspaceId),
    ),
  );

  return {
    tokenStorage,
    uuidGen,
    users,
    ingredients,
    authSrv: new AuthSrv(users, tokenStorage),
    ingredientSrv: new IngredientSrv(ingredients),
  };
};

import * as state from '../schema/state_pb';
import * as service from '../schema/service_pb';

function toInvoke(chaincodeId: string, fun: string): string {
  return `contract/testing/ccbio/invoke/${chaincodeId}:${fun}`;
}

function toQuery(chaincodeId: string, fun: string): string {
  return `contract/testing/ccbio/query/${chaincodeId}:${fun}`;
}

import { AxiosInstance } from 'axios';
import { JsonWriteStringOptions } from '@bufbuild/protobuf';

const submit_options: Partial<JsonWriteStringOptions> = {
  enumAsInteger: true,
  emitDefaultValues: true,
  useProtoFieldName: true,
};

export class CCBioAuthApi {
  /**
   *
   */
  constructor(api: AxiosInstance) {
    this.api = api;
  }

  api: AxiosInstance;
  chaincodeId = 'ccbio.Auth';

  async GetCurrentUser(): Promise<state.User> {
    const fn_name = 'GetCurrentUser';
    // const req_data = new service.GetCurrentUserRequest().toJsonString();
    const res = await this.api.post(toQuery(this.chaincodeId, fn_name));

    const res_data = res.data;
    const res_obj = state.User.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async GetUser(req: service.GetUserRequest): Promise<state.User> {
    const fn_name = 'GetUser';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);
    const res = await this.api.post(
      toQuery(this.chaincodeId, fn_name),
      req_data
    );



    const res_data = res.data;
    const res_obj = state.User.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async GetUserList(): Promise<state.User_List> {
    const fn_name = 'GetUserList';
    // const req_data = new service.GetUserListRequest().toJsonString();
    const res = await this.api.post(toQuery(this.chaincodeId, fn_name));

    const res_data = res.data;
    const res_obj = state.User_List.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async GetCollection(
    req: service.GetCollectionRequest
  ): Promise<state.Collection> {
    const fn_name = 'GetCollection';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toQuery(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;

    const res_obj = state.Collection.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async GetCollectionList(): Promise<state.Collection_List> {
    const fn_name = 'GetCollectionList';
    // const req_data = new service.GetCollectionListRequest().toJsonString();
    const res = await this.api.post(toQuery(this.chaincodeId, fn_name));

    const res_data = res.data;
    console.log(res_data);
    const res_obj = state.Collection_List.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async UserRegister(req: service.UserRegisterRequest): Promise<state.User> {
    const fn_name = 'UserRegister';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.User.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async CollectionCreate(
    req: service.CollectionCreateRequest
  ): Promise<state.Collection> {
    const fn_name = 'CollectionCreate';

    const t = req.collection?.accessControl?.toJson({ enumAsInteger: true });
    console.log(t);

    const j = req.toJson({
      enumAsInteger: true,
      emitDefaultValues: true,
    });
    console.log(j);

    const req_data = req.toJson(submit_options);

    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;

    const res_obj = state.Collection.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async CollectionUpdate(
    req: service.CollectionUpdateRequest
  ): Promise<state.Collection> {
    const fn_name = 'CollectionUpdate';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;

    const res_obj = state.Collection.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }
}

export class CCBioSpecimenApi {
  constructor(api: AxiosInstance) {
    this.api = api;
  }

  api: AxiosInstance;
  chaincodeId = 'ccbio.Specimen';

  async GetSpecimen(req: service.GetSpecimenRequest): Promise<state.Specimen> {
    const fn_name = 'GetSpecimen';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toQuery(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.Specimen.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async GetSpecimenList(): Promise<state.Specimen_List> {
    const fn_name = 'GetSpecimenList';
    // const req_data = new service.GetSpecimenListRequest().toJsonString();
    const res = await this.api.post(toQuery(this.chaincodeId, fn_name));

    const res_data = res.data;
    const res_obj = state.Specimen_List.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async GetSpecimenByCollection(
    req: service.GetSpecimenByCollectionRequest
  ): Promise<state.Specimen_List> {
    const fn_name = 'GetSpecimenByCollection';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);
    const res = await this.api.post(
      toQuery(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    console.log(res_data);
    if (!res_data) {
      return new state.Specimen_List();
    }

    // temporary fix
    const res_obj = new state.Specimen_List({
      items: res_data,
    });

    // const res_obj = state.Specimen_List.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async GetSpecimenHistory(
    req: service.GetSpecimenHistoryRequest
  ): Promise<state.Specimen_History> {
    const fn_name = 'GetSpecimenHistory';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);
    const res = await this.api.post(
      toQuery(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.Specimen_History.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async SpecimenCreate(
    req: service.SpecimenCreateRequest
  ): Promise<state.Specimen> {
    const fn_name = 'SpecimenCreate';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);
    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.Specimen.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async SpecimenUpdate(
    req: service.SpecimenUpdateRequest
  ): Promise<state.Specimen> {
    const fn_name = 'SpecimenUpdate';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);
    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.Specimen.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async SpecimenDelete(
    req: service.SpecimenDeleteRequest
  ): Promise<state.Specimen> {
    const fn_name = 'SpecimenDelete';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);
    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.Specimen.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: ${res_ob}`);
    return res_obj;
  }

  async SpecimenHideTransaction(
    req: service.SpecimenHideTxRequest
  ): Promise<state.Specimen> {
    const fn_name = 'SpecimenHideTransaction';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.Specimen.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async SpecimenUnHideTransaction(
    req: service.SpecimenUnHideTxRequest
  ): Promise<state.Specimen> {
    const fn_name = 'SpecimenUnHideTransaction';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.Specimen.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async SuggestedUpdateCreate(
    req: service.SuggestedUpdateCreateRequest
  ): Promise<state.SuggestedUpdate> {
    const fn_name = 'SuggestedUpdateCreate';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.SuggestedUpdate.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }

  async SpecimenUpdateApprove(
    req: service.SuggestedUpdateApproveRequest
  ): Promise<state.Specimen> {
    const fn_name = 'SpecimenUpdateApprove';
    const req_data = req.toJson(submit_options);
    console.log(`fn ${fn_name} => Response: ${req_data}`);

    const res = await this.api.post(
      toInvoke(this.chaincodeId, fn_name),
      req_data
    );

    const res_data = res.data;
    const res_obj = state.Specimen.fromJson(res_data);
    console.log(`fn ${fn_name} => Response: `);
    console.log(res_obj);
    return res_obj;
  }
}

export class CCBioApi {
  auth: CCBioAuthApi;
  specimen: CCBioSpecimenApi;

  constructor(api: AxiosInstance) {
    this.auth = new CCBioAuthApi(api);
    this.specimen = new CCBioSpecimenApi(api);
  }
}

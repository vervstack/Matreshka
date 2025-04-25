import {SortType} from "@vervstack/matreshka";

export type ListServicesReq = {
    paging: Paging;

    sort: Sort;
    searchPattern: string;
};

export type Paging = {
    limit: number;
    offset: number;
}
export type Sort = {
    type: SortType
    desc: boolean
}

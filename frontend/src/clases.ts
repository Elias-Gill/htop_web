// parseo de la api
export interface Nodo_api {
    Pid: Int32Array;
    Name: string;
    MemUsage: Float32Array;
}
export interface Data {
    Node: Nodo_api;
    Childs: Data[];
    Indent: number;
}


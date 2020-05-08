import { DumbRepository } from "../core/DumbRepository";

export class RestDumbRespository implements DumbRepository {
  async save(dumbName: string): Promise<void> {
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name: dumbName }),
    };

    await fetch("http://localhost:5000/api/dumb", requestOptions);

    return Promise.resolve();
  }
}

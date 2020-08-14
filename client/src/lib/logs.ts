export class Log {
    id: string;
    provider: string;
    raw_log: string;
    message: string;
    level: string;
    timestamp: Date;
    payload: object;
    preparedPayload: object | null;

    public constructor(id: string, provider: string, raw_log: string, message: string, level: string, timestamp: string, payload: object) {
        this.id = id;
        this.provider = provider;
        this.raw_log = raw_log;
        this.message = message;
        this.level = level;
        this.timestamp = new Date(timestamp);
        this.payload = payload;
        this.preparedPayload = null;
    }

    public getPayload() {
        if (this.preparedPayload === null) {
            this.preparedPayload = this.transformAttrDotToObject(this.payload);
        }

        return this.preparedPayload;
    }

    public getPayloadAttr(attr: string): string {
        if (this.payload === null) {
            return '';
        }

        return this.payload[attr] || '';
    }

    private transformAttrDotToObject(data: object): object {
        let transformed = {};

        Object.keys(data).forEach((attr) => {
            const val = data[attr]
            const keys = attr.split('.')

            if (keys.length === 1) {
                if (typeof val === "object") {
                    transformed[attr] = this.transformAttrDotToObject(val)
                } else {
                    if (val == this.message || val == this.level) {
                        return;
                    }

                    transformed[attr] = val;
                }
                return;
            }

            const k = keys.shift();
            const newKeys = keys.join('.');
            let nData = {};

            if (transformed.hasOwnProperty(k)) {
                nData[k] = transformed[k];
            } else {
                nData[k] = {};
            }

            nData[k][newKeys] = val;

            nData[k] = this.transformAttrDotToObject(nData[k]);

            transformed = {...transformed, ...nData};
        });

        return transformed;
    }
}
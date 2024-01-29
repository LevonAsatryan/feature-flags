import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import mongoose, { HydratedDocument } from 'mongoose';
import { KeyValue } from 'src/shared/types/shared-types';

@Schema()
export class Company {
  @Prop()
  name: string;

  @Prop({ unique: true, index: true, type: mongoose.Schema.Types.ObjectId })
  id: string;

  @Prop()
  featureFlags: KeyValue<string>[];
}

export type CompanyDocument = HydratedDocument<Company>;

export const CompanySchema = SchemaFactory.createForClass(Company);

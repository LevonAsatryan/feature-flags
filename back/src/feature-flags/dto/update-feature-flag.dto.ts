import { PartialType } from '@nestjs/mapped-types';
import { CreateFeatureFlagDto } from './create-feature-flag.dto';
import {
  IsBoolean,
  IsNumber,
  IsString,
  MaxLength,
  MinLength,
} from 'class-validator';

export class UpdateFeatureFlagDto extends PartialType(CreateFeatureFlagDto) {
  @IsString()
  @MaxLength(100)
  @MinLength(3)
  name?: string;

  @IsBoolean()
  isEnabled?: boolean;

  @IsNumber()
  containerId?: number;
}

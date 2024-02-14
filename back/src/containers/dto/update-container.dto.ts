import { PartialType } from '@nestjs/mapped-types';
import { CreateContainerDto } from './create-container.dto';
import { IsNumber, IsString, MaxLength, MinLength } from 'class-validator';

export class UpdateContainerDto extends PartialType(CreateContainerDto) {
  @IsString()
  @MaxLength(100)
  @MinLength(3)
  name?: string;

  @IsNumber()
  parentContainerId?: number;
}
